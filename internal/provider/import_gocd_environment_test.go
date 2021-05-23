package provider

import (
	"context"
	"fmt"
	"github.com/beamly/go-gocd/gocd"
	r "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

func testResourceEnvironmentImportBasic(t *testing.T) {
	suffix := randomString(10)
	rscId := "test-" + suffix
	resourceName := "gocd_environment." + rscId

	r.Test(t, r.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testGocdProviders,
		CheckDestroy: testGocdEnvironmentDestroy,
		Steps: []r.TestStep{
			{
				Config: testGocdEnvironmentConfig(suffix),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     rscId,
			},
		},
	})
}
func testGocdEnvironmentDestroy(s *terraform.State) error {

	gocdclient := testGocdProvider.Meta().(*gocd.Client)

	root := s.RootModule()
	for _, rs := range root.Resources {
		if rs.Type != "gocd_environment" {
			continue
		}

		name := rs.Primary.Attributes["name"]

		env, _, err := gocdclient.Environments.Get(context.Background(), name)
		if err == nil {
			return fmt.Errorf("still exists")
		}
		if env.Name == "" {
			return nil
		}
	}

	return nil
}

func testGocdEnvironmentConfig(suffix string) string {
	return strings.Replace(
		testFile("resource_environment.0.rsc.tf"),
		"test-environment",
		"test-"+suffix,
		-1,
	)
}
