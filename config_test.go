package config_test

import (
	"testing"

	config "github.com/mitz-it/golang-config"
	"github.com/stretchr/testify/assert"
)

func TestConfig_WhenPrefixAndPath_ShouldGetEnvVarWithoutPrefix(t *testing.T) {
	// arrange
	start := config.StartConfig{
		Prefix:     "MTZ",
		ConfigPath: "./sample.env",
	}

	cfg := config.NewConfig(start)

	// act
	dummy := cfg.Standard.GetString("dummy")

	// assert
	assert.Equal(t, "DUMMY", dummy)
}

func TestConfig_WhenOnlyPath_ShouldGetEnvVarWithPrefix(t *testing.T) {
	// arrange
	start := config.StartConfig{
		ConfigPath: "./sample.env",
	}

	cfg := config.NewConfig(start)
	// act

	dummy := cfg.Standard.GetString("mtz_dummy")

	// assert
	assert.Equal(t, "DUMMY", dummy)
}

func TestConfig_WhenOnlyPrefixAndNoEnv_ShouldGetEmptyVariable(t *testing.T) {
	// arrange
	start := config.StartConfig{
		Prefix: "MTZ",
	}

	cfg := config.NewConfig(start)

	// act
	dummy := cfg.Standard.GetString("dummy")

	// assert
	assert.Equal(t, "", dummy)
}
