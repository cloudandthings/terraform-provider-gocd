package gocd

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLogging(t *testing.T) {
	existingType := getLogType()
	existingLevel := getLogLevel()
	t.Run("Type", func(t *testing.T) {
		for _, lType := range []string{"JSON", "TEXT"} {
			os.Setenv(LogTypeEnvVarName, lType)
			assert.Equal(t, lType, getLogType())
		}
	})

	t.Run("Level", func(t *testing.T) {
		for _, lLevel := range []string{"PANIC", "FATAL", "ERROR", "WARNING", "INFO", "DEBUG"} {
			os.Setenv(LogLevelEnvVarName, lLevel)
			assert.Equal(t, lLevel, getLogLevel())
		}
	})

	os.Setenv(LogTypeEnvVarName, existingType)
	os.Setenv(LogLevelEnvVarName, existingLevel)
}
