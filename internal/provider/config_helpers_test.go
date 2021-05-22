package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestConfigHelper(t *testing.T) {
	t.Run("DecodeConfigStringList/SuccesSingle", decodeConfigStringListSuccessSingle)
	t.Run("DecodeConfigStringList/SuccesMulti", decodeConfigStringListSuccessMulti)
	t.Run("DecodeConfigStringList/FailInt", decodeConfigStringListFailInt)
	//t.Run("DefinitionDocFinish/Success", testDefinitionDocFinishSuccess)
	t.Run("DefinitionDocFinish/Fail", testDefinitionDocFinishFail)
	t.Run("RegexRuleSetValidator", testRegexRuleSetValidator)
	t.Run("SupressJsonDiff", testSupressJSONDiffs)
	t.Run("SupressJsonDiffPanic", testSupressJSONDiffsPanic)
}

func testRegexRuleSetValidator(t *testing.T) {
	matchErr := func(errs []error, r *regexp.Regexp) bool {
		// err must match one provided
		for _, err := range errs {
			if r.MatchString(err.Error()) {
				return true
			}
		}

		return false
	}

	for i, test := range []struct {
		f             schema.SchemaValidateFunc
		value         interface{}
		expectedError *regexp.Regexp
	}{
		{
			f: RegexRuleset(RegexRules{
				`^[a-zA-Z0-9_\-]{1}`: "first character of %q (%q) must be alphanumeric, underscore, or dot",
			}),
			value:         "$hallo-world",
			expectedError: regexp.MustCompile(`first character of "[^"]+" \("[^"]+"\) must be alphanumeric, underscore, or dot`),
		},
		{
			f: RegexRuleset(RegexRules{
				`^[a-zA-Z0-9_\-]{1}[a-zA-Z0-9_\-.]*$`: "only alphanumeric, underscores, hyphens, or dots allowed in %q (%q)",
			}),
			value:         "hallo-wo$rld",
			expectedError: regexp.MustCompile(`only alphanumeric, underscores, hyphens, or dots allowed in "[^"]+" \("[^"]+"\)`),
		},
	} {
		_, errs := test.f(test.value, "test_property")

		if test.expectedError == nil && len(errs) > 0 {
			continue
		}

		if len(errs) != 0 && test.expectedError == nil {
			t.Fatalf("expected test case %d to produce no errors, got %v", i, errs)
		}

		if !matchErr(errs, test.expectedError) {
			t.Fatalf("expected test case %d to produce error matching \"%s\", got %v", i, test.expectedError, errs)
		}

	}
}

func testDefinitionDocFinishFail(t *testing.T) {
	err := definitionDocFinish(
		&schema.ResourceData{},
		make(chan int),
	)
	assert.NotNil(t, err)
}

//func testDefinitionDocFinishSuccess(t *testing.T) {
//	expectedJSON := `{
//  "one": "hello",
//  "two": "world"
//}`
//	rd := (&schema.Resource{Schema: map[string]*schema.Schema{
//		"json": {Type: schema.TypeString, Computed: true},
//	}}).Data(&terraform.InstanceState{})
//	st := map[string]string{"one": "hello", "two": "world"}
//	err := definitionDocFinish(rd, st)
//
//	assert.Nil(t, err)
//	assert.Equal(t, expectedJSON, rd.Get("json"))
//	assert.Equal(t, "3710939758", rd.Id())
//}

func decodeConfigStringListFailInt(t *testing.T) {
	n := []int{6, 7, 8}
	i := make([]interface{}, len(n))
	for j := range n {
		i[j] = n[j]
	}
	assert.Panics(t, func() { decodeConfigStringList(i) })
}

func decodeConfigStringListSuccessSingle(t *testing.T) {
	s := []string{"one"}
	i := make([]interface{}, len(s))
	for j := range s {
		i[j] = s[j]
	}
	strs := decodeConfigStringList(i)

	assert.Len(t, strs, 1)
	assert.Equal(t, strs[0], "one")
}

func decodeConfigStringListSuccessMulti(t *testing.T) {
	s := []string{"one", "two"}
	i := make([]interface{}, len(s))
	for j := range s {
		i[j] = s[j]
	}
	strs := decodeConfigStringList(i)

	assert.Len(t, strs, 2)
	assert.Equal(t, strs[0], "one")
	assert.Equal(t, strs[1], "two")
}

func testSupressJSONDiffsPanic(t *testing.T) {
	type args struct {
		old string
		new string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "EmptyNewOld",
			args: args{
				new: "",
				old: "",
			},
		},
		{
			name: "BadOld",
			args: args{
				old: "{ items: [ { id: 1, name: 'test1' }, { id: 2, name: 'test2' } ] }",
				new: "{\"hallo\": \"world\"}",
			},
		},
		{
			name: "BadNew",
			args: args{
				old: "{\"hallo\": \"world\"}",
				new: "{ items: [ { id: 1, name: 'test1' }, { id: 2, name: 'test2' } ] }",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Panics(t, func() { supressJSONDiffs("", tt.args.old, tt.args.new, nil) })
		})
	}
}

func testSupressJSONDiffs(t *testing.T) {
	type args struct {
		old string
		new string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "EmptyOld",
			args: args{old: "", new: "mock"},
			want: false,
		},
		{
			name: "EmptyNew",
			args: args{old: "mock", new: ""},
			want: false,
		},
		{
			name: "Equal",
			args: args{old: "{\"hallo\": \"world\"}", new: "{\"hallo\": \"world\"}"},
			want: true,
		},
		{
			name: "Unequal",
			args: args{old: "{\"hallo\": \"world\"}", new: "{\"hallo\": \"foobar\"}"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := supressJSONDiffs("", tt.args.old, tt.args.new, nil)
			assert.Equal(t, got, tt.want)
		})
	}
}
