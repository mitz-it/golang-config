package config_test

import (
	"fmt"
	"testing"
	"time"

	config "github.com/mitz-it/golang-config"
	v "github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLoadEnv_WhenUseCorrectPrefixAndPath_ShouldInstanciateModuleInstance(t *testing.T) {
	// arrange
	config.LoadEnv("MTZ", "env/dummy.env")

	// act
	check := func() {
		fmt.Printf("%t", config.Env != nil)
	}

	// assert
	assert.NotPanics(t, check)
}

func TestEnv_WhenAccessingNotLoadedModuleInstance_ShouldPanic(t *testing.T) {
	// arrange
	config.Env = nil

	// act
	check := func() {
		fmt.Printf("%s", config.Env.GetString("dummy"))
	}

	// assert
	assert.Panics(t, check)
}

func TestGet_WhenGivenTheRightKey_ShouldReturnInterface(t *testing.T) {
	// arrange
	config.LoadEnv("MTZ", "env/dummy.env")

	// act
	i := config.Env.Get("interface")

	// assert
	assert.Equal(t, "20", i)
}

func TestGetString_WhenGivenTheRightKey_ShouldReturnString(t *testing.T) {
	// arrange
	config.LoadEnv("MTZ", "env/dummy.env")

	// act
	dummy := config.Env.Get("dummy")

	// assert
	assert.IsType(t, "DUMMY", dummy)
}

func TestGetStringSlice_WhenGivenTheRightKey_ShouldReturnStringSlice(t *testing.T) {
	// arrange
	config.LoadEnv("MTZ", "env/dummy.env")

	// act
	slice := config.Env.GetStringSlice("slice")

	// assert
	assert.Equal(t, []string{"val1", "val2"}, slice)
}

func TestGetStringMap_WhenGivenTheRightKey_ShouldReturnStringMap(t *testing.T) {
	// arrange
	config.LoadEnv("MTZ", "env/dummy.env")

	expected := map[string]interface{}{
		"key1": "val1",
		"key2": "val2",
	}

	// act
	m := config.Env.GetStringMap("map")

	// assert
	assert.Equal(t, expected, m)
}

func TestGetBool_WhenGivenTheRightKey_ShouldReturnBool(t *testing.T) {
	// arrange
	config.LoadEnv("MTZ", "env/dummy.env")

	// act
	b := config.Env.GetBool("bool")

	// assert
	assert.True(t, b)
}

func TestGetInt_WhenGivenTheRightKey_ShouldReturnInt(t *testing.T) {
	// arrange
	config.LoadEnv("MTZ", "env/dummy.env")

	// act
	integer := config.Env.GetInt("int")

	// assert
	assert.Equal(t, 1, integer)
}

func TestGetFloat64_WhenGivenTheRightKey_ShouldReturnFloat64(t *testing.T) {
	// arrange
	config.LoadEnv("MTZ", "env/dummy.env")
	expected := 4.66

	// act
	integers := config.Env.GetFloat64("float")

	// assert
	assert.Equal(t, expected, integers)
}

func TestGetTime_WhenGivenTheRightKey_ShouldReturnTime(t *testing.T) {
	// arrange
	config.LoadEnv("MTZ", "env/dummy.env")
	expected, _ := time.Parse(time.RFC3339, "2020-04-03T22:50:45Z")

	// act
	date := config.Env.GetTime("time")

	// assert
	assert.Equal(t, expected, date)
}

func TestGetDuration_WhenGivenTheRightKey_ShouldReturnDuration(t *testing.T) {
	// arrange
	config.LoadEnv("MTZ", "env/dummy.env")
	expected, _ := time.ParseDuration("4h")

	// act
	hours := config.Env.GetDuration("duration")

	// assert
	assert.Equal(t, expected, hours)
}

func TestGetViper_WhenCallingViperFunction_ShouldReturnInstanceFromViper(t *testing.T) {
	// arrange
	config.LoadEnv("MTZ", "env/dummy.env")
	expected := v.New()

	// act
	viper := config.Env.Viper()

	// assert
	assert.IsType(t, expected, viper)
}
