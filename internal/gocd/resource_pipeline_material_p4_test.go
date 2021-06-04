package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testUnmarshallMaterialAttributesP4(t *testing.T) {
	m := MaterialAttributesP4{}
	expected := MaterialAttributesP4{
		Name:       "test-name",
		Port:       "test-port",
		UseTickets: true,
		View:       "test-view",

		Username:          "test-user",
		Password:          "test-pass",
		EncryptedPassword: "test-encryptedpass",

		Destination: "test-dest",
		Filter: &MaterialFilter{
			Ignore: []string{"one", "two"},
		},
		InvertFilter: true,
		AutoUpdate:   true,
	}
	unmarshallMaterialAttributesP4(&m, map[string]interface{}{
		"name":               expected.Name,
		"port":               expected.Port,
		"use_tickets":        expected.UseTickets,
		"view":               expected.View,
		"username":           expected.Username,
		"password":           expected.Password,
		"encrypted_password": expected.EncryptedPassword,
		"destination":        expected.Destination,
		"filter": map[string]interface{}{
			"ignore": expected.Filter.Ignore,
		},
		"auto_update":   expected.AutoUpdate,
		"invert_filter": expected.InvertFilter,
		"foo":           nil,
	})

	assert.Equal(t, expected, m)
}
