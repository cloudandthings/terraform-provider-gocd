package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigRepoGetVersion(t *testing.T) {
	repo := ConfigRepo{Version: "test-version"}
	version := repo.GetVersion()
	assert.Equal(t, repo.Version, version)
}

func TestConfigRepoSetVersion(t *testing.T) {
	repo := ConfigRepo{}
	assert.Equal(t, repo.Version, "")
	repo.SetVersion("test-version")
	assert.Equal(t, repo.Version, "test-version")
}
