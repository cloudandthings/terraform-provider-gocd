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

func testResourceEnvironmentAssociationImportBasic(t *testing.T) {
	suffix := randomString(10)
	rscId := "test-" + suffix
	resourceName := "gocd_environment_association." + rscId

	r.Test(t, r.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testGocdProviders,
		CheckDestroy: testGocdEnvironmentAssociationDestroy,
		Steps: []r.TestStep{
			{
				Config: testGocdEnvironmentAssociationConfig(suffix),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     "test-environment/p/" + rscId,
			},
		},
	})
}
func testGocdEnvironmentAssociationDestroy(s *terraform.State) error {

	gocdclient := testGocdProvider.Meta().(*gocd.Client)

	root := s.RootModule()
	for _, rs := range root.Resources {
		if rs.Type != "gocd_environment_association" {
			continue
		}

		name := rs.Primary.Attributes["environment"]

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

func testGocdEnvironmentAssociationConfig(suffix string) string {
	return strings.Replace(
		testFile("resource_environment_association.0.rsc.tf"),
		"test-pipeline", // This is not a type, it should say "test-pipeline" as the name of the associaiton is kind of irrelevant.
		"test-"+suffix,
		-1,
	)
}
