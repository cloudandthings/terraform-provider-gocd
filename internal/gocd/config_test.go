package gocd

import (
	"github.com/stretchr/testify/assert"
	"os"
	"os/user"
	"testing"
)

func TestConfig(t *testing.T) {
	path, err := ConfigFilePath()
	u, err := user.Current()

	assert.Nil(t, err)
	assert.Equal(t, u.HomeDir+"/.gocd.conf", path)
}

func TestConfigCustom(t *testing.T) {
	os.Setenv("GOCD_CONFIG_PATH", "/mock/path/gocd.conf")
	path, err := ConfigFilePath()

	assert.Nil(t, err)
	assert.Equal(t, "/mock/path/gocd.conf", path)
}
