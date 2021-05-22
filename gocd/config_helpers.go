package gocd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"hash/crc32"
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
	d.SetId(strconv.Itoa(hascode{}.String(jsonString)))

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

// https://www.terraform.io/docs/extend/guides/v2-upgrade-guide.html#removal-of-helper-hashcode-package
type hascode struct {}

// String hashes a string to a unique hashcode.
//
// crc32 returns a uint32, but for our use we need
// and non negative integer. Here we cast to an integer
// and invert it if the result is negative.
func ( h hascode) String(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

// Strings hashes a list of strings to a unique hashcode.
func ( h hascode) Strings(strings []string) string {
	var buf bytes.Buffer

	for _, s := range strings {
		buf.WriteString(fmt.Sprintf("%s-", s))
	}

	return fmt.Sprintf("%d", h.String(buf.String()))
}