package gocd

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestConfigRepo(t *testing.T) {
	setup()
	defer teardown()

	t.Run("List", testConfigRepoList)
	t.Run("Get", testConfigRepoGet)
	t.Run("Create", testConfigRepoCreate)
	t.Run("Update", testConfigRepoUpdate)
	t.Run("Delete", testConfigRepoDelete)
}

func testConfigRepoList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/admin/config_repos", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		j, _ := ioutil.ReadFile("test/resources/configrepos.1.json")
		fmt.Fprint(w, string(j))
	})

	repos, _, err := client.ConfigRepos.List(context.Background())

	assert.Nil(t, err)
	assert.Len(t, repos, 1)

	testConfigRepo(t, repos[0])
}

func testConfigRepoGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/admin/config_repos/repo1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		j, _ := ioutil.ReadFile("test/resources/configrepos.0.json")
		w.Header().Set("Etag", "mock-etag")
		fmt.Fprint(w, string(j))
	})

	repo, _, err := client.ConfigRepos.Get(context.Background(), "repo1")

	assert.Nil(t, err)
	assert.Equal(t, "mock-etag", repo.Version)
	testConfigRepo(t, repo)
}

func testConfigRepoCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/admin/config_repos", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Unexpected HTTP method")
		j, _ := ioutil.ReadFile("test/resources/configrepos.0.json")
		fmt.Fprint(w, string(j))
	})

	r := ConfigRepo{ID: "repo1", PluginID: "json.config.plugin", Material: Material{Type: "git", Attributes: &MaterialAttributesGit{URL: "https://github.com/config-repo/gocd-json-config-example.git", Branch: "master", AutoUpdate: true}}}
	repo, _, err := client.ConfigRepos.Create(context.Background(), &r)

	assert.Nil(t, err)
	assert.NotNil(t, repo)
	testConfigRepo(t, repo)
}

func testConfigRepoUpdate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/admin/config_repos/repo1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method, "Unexpected HTTP method")
		j, _ := ioutil.ReadFile("test/resources/configrepos.0.json")
		assert.Equal(t, `"test-version"`, r.Header.Get("If-Match"))
		w.Header().Set("ETag", `"mock-version"`)
		fmt.Fprint(w, string(j))
	})

	r := ConfigRepo{ID: "repo1", PluginID: "json.config.plugin", Material: Material{Type: "git", Attributes: &MaterialAttributesGit{URL: "https://github.com/config-repo/gocd-json-config-example.git", Branch: "master", AutoUpdate: true}}, Version: "test-version"}
	repo, _, err := client.ConfigRepos.Update(context.Background(), "repo1", &r)

	assert.Nil(t, err)
	assert.NotNil(t, repo)
	testConfigRepo(t, repo)
	assert.Equal(t, repo.Version, "mock-version")
}

func testConfigRepoDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/admin/config_repos/repo1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "DELETE", "Unexpected HTTP method")
		assert.Equal(t, r.Header.Get("Accept"), apiV1)
		fmt.Fprint(w, `{
										  "message": "The config repo 'repo1' was deleted successfully."
										}`)
	})
	message, resp, err := client.ConfigRepos.Delete(context.Background(), "repo1")

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "The config repo 'repo1' was deleted successfully.", message)
}

func testConfigRepo(t *testing.T, repo *ConfigRepo) {

	for _, attribute := range []EqualityTest{
		{repo.Links.Get("Self").URL.String(), "https://ci.example.com/go/api/admin/config_repos/repo1"},
		{repo.Links.Get("Doc").URL.String(), "https://api.gocd.org/#config-repos"},
		{repo.Links.Get("Find").URL.String(), "https://ci.example.com/go/api/admin/config_repos/:id"},
		{repo.ID, "repo1"},
		{repo.PluginID, "json.config.plugin"},
		{repo.Material.Type, "git"},
		{repo.Material.Attributes.(*MaterialAttributesGit).URL, "https://github.com/config-repo/gocd-json-config-example.git"},
		{repo.Material.Attributes.(*MaterialAttributesGit).Branch, "master"},
	} {
		assert.Equal(t, attribute.wanted, attribute.got)
	}
}
