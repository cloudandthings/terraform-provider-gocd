package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testUnmarshallMaterialAttributesDependency(t *testing.T) {
	m := MaterialAttributesDependency{}
	unmarshallMaterialAttributesDependency(&m, map[string]interface{}{
		"name":        "test-name",
		"pipeline":    "test-pipeline",
		"stage":       "test-stage",
		"foo":         nil,
		"auto_update": true,
	})

	assert.Equal(t, MaterialAttributesDependency{
		Name:       "test-name",
		Pipeline:   "test-pipeline",
		Stage:      "test-stage",
		AutoUpdate: true,
	}, m)
}
