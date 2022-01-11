package gocd

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestAuthentication(t *testing.T) {
	setup()
	defer teardown()

	t.Run("Login", testAuthenticationLogin)
	t.Run("LoginFail", testAuthenticationLoginFail)
}

func testAuthenticationLoginFail(t *testing.T) {
	env := os.Getenv("GOCD_RAISE_ERROR_NEW_REQUEST")
	os.Setenv("GOCD_RAISE_ERROR_NEW_REQUEST", "yes")

	err := client.Login(context.Background())
	assert.EqualError(t, err, "Mock Testing Error")

	os.Setenv("GOCD_RAISE_ERROR_NEW_REQUEST", env)
}

func testAuthenticationLogin(t *testing.T) {

	mockCookie := "JSESSIONID=hash;Path=/go;Expires=Mon, 15-Jun-2015 10:16:20 GMT"

	mux.HandleFunc("/api/api/agents", func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.Header["Accept"], "application/vnd.go.cd.v2+json")
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		testAuth(t, r, mockAuthorization)

		w.Header().Set("Set-Cookie", mockCookie)

		j, _ := ioutil.ReadFile("test/resources/agents.0.json")
		fmt.Fprint(w, string(j))
	})

	client.Login(context.Background())

	assert.Equal(t, client.cookie, mockCookie)

}
