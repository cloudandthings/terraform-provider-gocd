package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testUnmarshallMaterialAttributesTfs(t *testing.T) {
	m := MaterialAttributesTfs{}
	expected := MaterialAttributesTfs{
		Name: "test-name",

		URL:         "test-url",
		ProjectPath: "test-project-path",
		Domain:      "test-domain",

		Username:          "test-username",
		Password:          "test-password",
		EncryptedPassword: "test-encrypted-password",

		Destination: "test-destination",
		Filter: &MaterialFilter{
			Ignore: []string{"one", "two"},
		},
		InvertFilter: true,
		AutoUpdate:   true,
	}
	unmarshallMaterialAttributesTfs(&m, map[string]interface{}{
		"name":               expected.Name,
		"url":                expected.URL,
		"project_path":       expected.ProjectPath,
		"domain":             expected.Domain,
		"username":           expected.Username,
		"password":           expected.Password,
		"encrypted_password": expected.EncryptedPassword,
		"destination":        expected.Destination,
		"invert_filter":      expected.InvertFilter,
		"filter": map[string]interface{}{
			"ignore": expected.Filter.Ignore,
		},
		"auto_update": expected.AutoUpdate,
		"foo":         nil,
	})

	assert.Equal(t, expected, m)
}
