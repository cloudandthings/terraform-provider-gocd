package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testUnmarshallMaterialAttributesSvn(t *testing.T) {
	m := MaterialAttributesSvn{}
	expected := MaterialAttributesSvn{
		Name:              "test-name",
		URL:               "test-url",
		Username:          "test-username",
		Password:          "test-password",
		EncryptedPassword: "test-encrypted-password",

		CheckExternals: true,

		Destination: "test-destination",
		Filter: &MaterialFilter{
			Ignore: []string{"one", "two"},
		},
		InvertFilter: true,
		AutoUpdate:   true,
	}
	unmarshallMaterialAttributesSvn(&m, map[string]interface{}{
		"name":               expected.Name,
		"url":                expected.URL,
		"username":           expected.Username,
		"password":           expected.Password,
		"encrypted_password": expected.EncryptedPassword,
		"check_externals":    expected.CheckExternals,
		"destination":        expected.Destination,
		"filter": map[string]interface{}{
			"ignore": expected.Filter.Ignore,
		},
		"invert_filter": expected.InvertFilter,
		"auto_update":   true,
		"foo":           nil,
	})

	assert.Equal(t, expected, m)
}
