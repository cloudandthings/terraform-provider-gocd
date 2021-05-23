package provider

import (
	"fmt"
	r "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func testDataSourceJobDefinition(t *testing.T) {
	for i := 0; i <= 1; i++ {
		t.Run(
			fmt.Sprintf("gocd_job_definition.%d", i),
			dataSourceJobDefinition(t, i,
				fmt.Sprintf("data_source_job_definition.%d.rsc.tf", i),
				fmt.Sprintf("data_source_job_definition.%d.rsp.json", i),
			),
		)
	}

}

func dataSourceJobDefinition(_ *testing.T, index int, configPath string, expectedPath string) func(t *testing.T) {
	return func(t *testing.T) {
		config := testFile(configPath)
		expected := testFile(expectedPath)
		r.UnitTest(t, r.TestCase{
			Providers: testGocdProviders,
			Steps: testStepComparisonCheck(&TestStepJSONComparison{
				Index:        index,
				ID:           "data.gocd_job_definition.test",
				Config:       config,
				ExpectedJSON: expected,
			}),
		})
	}
}
