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

func testResourcePipelineTemplateImportBasic(t *testing.T) {
	suffix := randomString(10)
	resourceName := fmt.Sprintf("gocd_pipeline_template.test-%s", suffix)

	r.Test(t, r.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testGocdProviders,
		CheckDestroy: testGocdPipelineTemplateDestroy,
		Steps: []r.TestStep{
			{
				Config: testGocdPipelineTemplateConfig(suffix),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testGocdPipelineTemplateDestroy(s *terraform.State) error {

	gocdclient := testGocdProvider.Meta().(*gocd.Client)

	root := s.RootModule()
	for _, rs := range root.Resources {
		if rs.Type != "gocd_pipeline_template" {
			continue
		}

		_, _, err := gocdclient.PipelineTemplates.Get(context.Background(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("still exists")
		}
	}

	return nil
}

func testGocdPipelineTemplateConfig(suffix string) string {
	return strings.Replace(
		testFile("resource_pipeline_template.0.rsc.tf"),
		"test-pipeline",
		"test-"+suffix,
		-1,
	)
}
