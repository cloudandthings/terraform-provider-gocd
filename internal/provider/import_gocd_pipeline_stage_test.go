package provider

import (
	"context"
	"fmt"
	"github.com/cloudandthings/terraform-provider-gocd/internal/gocd"
	r "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

func testResourcePipelineStageImportBasic(t *testing.T) {
	suffix := randomString(10)
	rscId := "test-" + suffix
	resourceName := "gocd_pipeline_stage." + rscId

	r.Test(t, r.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testGocdProviders,
		CheckDestroy: testGocdPipelineStageDestroy,
		Steps: []r.TestStep{
			{
				Config: testGocdPipelineStageConfig(suffix),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     "template/test-pipeline-template/" + rscId,
			},
		},
	})
}

func testGocdPipelineStageDestroy(s *terraform.State) error {

	gocdclient := testGocdProvider.Meta().(*gocd.Client)

	root := s.RootModule()
	for _, rs := range root.Resources {
		if rs.Type != "gocd_pipeline_stage" {
			continue
		}

		ptName := rs.Primary.Attributes["pipeline_template"]
		name := rs.Primary.Attributes["name"]

		pt, _, err := gocdclient.PipelineTemplates.Get(context.Background(), ptName)
		for _, stage := range pt.Stages {
			if stage.Name == name {
				return fmt.Errorf("still exists")
			}
		}
		if err == nil {
			return fmt.Errorf("still exists")
		}
	}

	return nil
}

func testGocdPipelineStageConfig(suffix string) string {
	return strings.Replace(
		testFile("resource_pipeline_stage.0.rsc.tf"),
		"test-stage",
		"test-"+suffix,
		-1,
	)
}
