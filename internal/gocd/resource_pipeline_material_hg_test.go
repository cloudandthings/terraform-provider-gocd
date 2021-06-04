package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testUnmarshallMaterialAttributesHg(t *testing.T) {
	m := MaterialAttributesHg{}
	expected := MaterialAttributesHg{
		Name: "test-name",
		URL:  "test-url",

		Destination:  "test-destination",
		InvertFilter: true,
		Filter: &MaterialFilter{
			Ignore: []string{"one", "two"},
		},
		AutoUpdate: true,
	}
	unmarshallMaterialAttributesHg(&m, map[string]interface{}{
		"name":          expected.Name,
		"url":           expected.URL,
		"destination":   expected.Destination,
		"invert_filter": expected.InvertFilter,
		"filter": map[string]interface{}{
			"ignore": expected.Filter.Ignore,
		},
		"auto_update": expected.AutoUpdate,
		"foo":         nil,
	})

	assert.Equal(t, expected, m)
}
