package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testUnmarshallMaterialAttributesPlugin(t *testing.T) {
	m := MaterialAttributesPlugin{}
	expected := MaterialAttributesPlugin{
		Ref: "test-ref",

		Destination: "test-desintation",
		Filter: &MaterialFilter{
			Ignore: []string{"one", "two"},
		},
		InvertFilter: true,
	}

	unmarshallMaterialAttributesPlugin(&m, map[string]interface{}{
		"ref":           expected.Ref,
		"destination":   expected.Destination,
		"invert_filter": expected.InvertFilter,
		"filter": map[string]interface{}{
			"ignore": expected.Filter.Ignore,
		},
		"foo": nil,
	})

	assert.Equal(t, expected, m)
}
