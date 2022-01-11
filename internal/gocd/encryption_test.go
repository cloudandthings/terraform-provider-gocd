package gocd

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestEncryption(t *testing.T) {
	setup()
	defer teardown()

	t.Run("BasicEncryption", testEncryptionEncrypt)
}

func testEncryptionEncrypt(t *testing.T) {
	mux.HandleFunc("/api/admin/encrypt", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Unexpected HTTP method")
		assert.Equal(t, apiV1, r.Header.Get("Accept"))
		var bdy []byte
		bdy, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, "{\n  \"value\": \"test-plaintext\"\n}\n", string(bdy))
		fmt.Fprint(w, `{
  "_links": {
    "self": {
      "href": "http://ci.example.com/go/api/admin/encrypt"
    },
    "doc": {
      "href": "https://api.gocd.org/#encryption"
    }
  },
  "encrypted_value": "mock-ciphertext"
}`)
	})
	ct, _, err := client.Encryption.Encrypt(context.Background(), "test-plaintext")
	if err != nil {
		t.Error(t, err)
	}

	assert.Equal(t, "mock-ciphertext", ct.EncryptedValue)
}
