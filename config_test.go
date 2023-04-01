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
		ConfigPath: "env/sample.env",
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
		ConfigPath: "env/sample.env",
	}

	cfg := config.NewConfig(start)
	// act

	dummy := cfg.Standard.GetString("mtz_dummy")

	// assert
	assert.Equal(t, "DUMMY", dummy)
}

func TestConfig_WhenEnvFileNotExists_ShouldPanic(t *testing.T) {
	// arrange
	start := config.StartConfig{
		Prefix:     "MTZ",
		ConfigPath: "env/invalid.env",
	}

	// act
	cfg := func() {
		config.NewConfig(start)
	}

	// assert
	assert.Panics(t, cfg)
}
