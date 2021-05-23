package provider

import (
	"encoding/json"
	"fmt"
	"github.com/cloudandthings/terraform-provider-gocd/internal/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"reflect"
	"regexp"
	"strconv"
)

// Give an abstract list of strings cast as []interface{}, convert them back to []string{}.
func decodeConfigStringList(lI []interface{}) []string {

	if len(lI) == 1 {
		return []string{lI[0].(string)}
	}
	ret := make([]string, len(lI))
	for i, vI := range lI {
		ret[i] = vI.(string)
	}
	return ret
}

// Take our object we parsed from the TF resource, and encode it in JSON.
func definitionDocFinish(d *schema.ResourceData, r interface{}) error {
	jsonDoc, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}
	jsonString := string(jsonDoc)
	d.Set("json", jsonString)
	d.SetId(strconv.Itoa(hashcode.String(jsonString)))

	return nil
}

func supressJSONDiffs(k, old, new string, d *schema.ResourceData) bool {
	var j1, j2 interface{}
	if (old == "" || new == "") && old != new {
		return false
	}
	if err := json.Unmarshal([]byte(old), &j1); err != nil {
		panic(err)
	}
	if err := json.Unmarshal([]byte(new), &j2); err != nil {
		panic(err)
	}
	return reflect.DeepEqual(j2, j1)
}

type RegexRules map[string]string

// RegexRuleset returns a SchemaValidateFunc which tests if all the provided regex rules
// successfully match the supplied string. Having multiple rule/reason
func RegexRuleset(rules RegexRules) schema.SchemaValidateFunc {
	return func(i interface{}, key string) (s []string, errors []error) {
		value, ok := i.(string)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %q (%q) to be string", key, value))
			return
		}

		for rule, reason := range rules {
			if !regexp.MustCompile(rule).MatchString(value) {
				errors = append(errors, fmt.Errorf(reason, key, value))
			}
		}

		return
	}
}
