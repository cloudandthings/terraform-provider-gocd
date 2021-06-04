package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testUnmarshallMaterialAttributesPkg(t *testing.T) {
	m := MaterialAttributesPackage{}
	expected := MaterialAttributesPackage{Ref: "test-ref"}

	unmarshallMaterialAttributesPackage(&m, map[string]interface{}{"ref": expected.Ref, "foo": nil})

	assert.Equal(t, expected, m)
}
