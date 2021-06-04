package provider

import (
	"fmt"
	"github.com/cloudandthings/terraform-provider-gocd/internal/gocd"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/sergi/go-diff/diffmatchpatch"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"
)

var (
	testGocdProviders map[string]*schema.Provider
	testGocdProvider  *schema.Provider
	testGocdClient    *gocd.Client
)

type TestStepJSONComparison struct {
	Index        int
	ID           string
	Config       string
	ExpectedJSON string
}

func init() {

	testGocdProvider = New("dev")()
	testGocdProviders = map[string]*schema.Provider{
		"gocd": testGocdProvider,
	}

	cfg := gocd.Configuration{
		Server:   os.Getenv("GOCD_URL"),
		Username: os.Getenv("GOCD_USERNAME"),
		Password: os.Getenv("GOCD_PASSWORD"),
	}

	testGocdClient = cfg.Client()
}

func TestProvider(t *testing.T) {
	if err := New("dev")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testStepComparisonCheck(t *TestStepJSONComparison) []resource.TestStep {
	return []resource.TestStep{
		{
			Config: t.Config,
			Check: func(s *terraform.State) error {
				root := s.RootModule()
				rs, ok := root.Resources[t.ID]
				if !ok {
					return fmt.Errorf("In '%d'.\nNot found: %s", t.Index, t.ID)
				}
				if rs.Primary.ID == "" {
					return fmt.Errorf("In '%d'.\nNo ID is set", t.Index)
				}

				if v := rs.Primary.Attributes["json"]; v != t.ExpectedJSON {
					dmp := diffmatchpatch.New()
					rawDiffs := dmp.DiffMain(v, t.ExpectedJSON, true)
					rawDiff := dmp.DiffPrettyText(rawDiffs)

					err := fmt.Errorf("In '%d'.\nValue mismatch for 'json' is:\n'%s'", t.Index, rawDiff)
					return err
				}

				return nil
			},
		},
	}
}

func testFile(name string) string {
	f, err := ioutil.ReadFile(fmt.Sprintf("test/%s", name))
	if err != nil {
		panic(err)
	}

	return string(f)
}

func testAccPreCheck(t *testing.T) {

	if u := os.Getenv("GOCD_URL"); u == "" {
		t.Fatal("GOCD_URL must be set for acceptance tests.")
	}
}

func randomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
