package gocd

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestConfiguration(t *testing.T) {
	setup()
	defer teardown()

	t.Run("HasAuth", testConfigurationHasAuth)
	t.Run("New", testConfigurationNew)
	t.Run("GetVersion", testConfigurationGetVersion)
	t.Run("Get", testConfigurationGet)
}

func testConfigurationGet(t *testing.T) {
	mux.HandleFunc("/api/admin/config.xml", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		j, _ := ioutil.ReadFile("test/resources/config.0.xml")
		fmt.Fprint(w, string(j))

		w.Header().Set("X-CRUISE-CONFIG-MD5", "c21b6c9f1b24a816cddf457548a987a9")
		w.Header().Set("Content-Type", "text/xml")
	})
	cfg, _, err := client.Configuration.Get(context.Background())
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, cfg.Server.ServerID, "9C0C0282-A554-457D-A0F8-9CF8A754B4AB")
	assert.Equal(t, cfg.Server.Security.PasswordFile.Path, "/etc/go/password.properties")
	assert.Equal(t, "defaultGroup", cfg.PipelineGroups[0].Name)
	assert.Len(t, cfg.PipelineGroups, 1)
}

func testConfigurationGetVersion(t *testing.T) {
	mux.HandleFunc("/api/version", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		j, _ := ioutil.ReadFile("test/resources/version.0.json")
		fmt.Fprint(w, string(j))
	})
	v, _, err := client.Configuration.GetVersion(context.Background())
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "https://build.go.cd/go/api/version", v.Links.Get("Self").URL.String())
	assert.Equal(t, "https://api.gocd.org/#version", v.Links.Get("Doc").URL.String())
	assert.Equal(t, "16.6.0", v.Version)
	assert.Equal(t, "3348", v.BuildNumber)
	assert.Equal(t, "a7a5717cbd60c30006314fb8dd529796c93adaf0", v.GitSHA)
	assert.Equal(t, "16.6.0 (3348-a7a5717cbd60c30006314fb8dd529796c93adaf0)", v.FullVersion)
	assert.Equal(t, "https://github.com/gocd/gocd/commits/a7a5717cbd60c30006314fb8dd529796c93adaf0", v.CommitURL)
}

func testConfigurationNew(t *testing.T) {
	c := Configuration{}
	client := c.Client()
	assert.NotNil(t, client)
}

func testConfigurationHasAuth(t *testing.T) {
	c := Configuration{}

	c.Username = "user"
	c.Password = "pass"
	assert.True(t, c.HasAuth())

	c.Username = "user"
	c.Password = ""
	assert.False(t, c.HasAuth())

	c.Username = ""
	c.Password = "pass"
	assert.False(t, c.HasAuth())

	c.Username = ""
	c.Password = ""
	assert.False(t, c.HasAuth())
}
