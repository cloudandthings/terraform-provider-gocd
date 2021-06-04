package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testUnmarshallMaterialAttributesGit(t *testing.T) {
	expected := MaterialAttributesGit{
		Name:   "test-name",
		URL:    "test-url",
		Branch: "test-branch",

		SubmoduleFolder: "test-submodule_folder",
		ShallowClone:    true,

		Destination: "test-destination",
		Filter: &MaterialFilter{
			Ignore: []string{"one", "two"},
		},
		InvertFilter: true,
		AutoUpdate:   true,
	}

	m := MaterialAttributesGit{}
	unmarshallMaterialAttributesGit(&m, map[string]interface{}{
		"name":             "test-name",
		"url":              expected.URL,
		"auto_update":      expected.AutoUpdate,
		"branch":           expected.Branch,
		"submodule_folder": expected.SubmoduleFolder,
		"destination":      expected.Destination,
		"shallow_clone":    expected.ShallowClone,
		"invert_filter":    expected.InvertFilter,
		"filter": map[string]interface{}{
			"ignore": expected.Filter.Ignore,
		},
		"foo": nil,
	})

	assert.Equal(t, expected, m)
}
