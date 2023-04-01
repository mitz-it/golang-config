package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	v "github.com/spf13/viper"
)

type Environment struct {
	viper *v.Viper
}

var Env *Environment

func LoadEnv(prefix, path string) *Environment {
	loadDotEnv(path)

	v := createViper(prefix)

	Env = &Environment{
		viper: v,
	}

	return Env
}

func (env *Environment) Get(key string) interface{} {
	return env.viper.Get(key)
}

func (env *Environment) GetString(key string) string {
	return env.viper.GetString(key)
}

func (env *Environment) GetStringSlice(key string) []string {
	return env.viper.GetStringSlice(key)
}

func (env *Environment) GetStringMap(key string) map[string]interface{} {
	return env.viper.GetStringMap(key)
}

func (env *Environment) GetBool(key string) bool {
	return env.viper.GetBool(key)
}

func (env *Environment) GetInt(key string) int {
	return env.viper.GetInt(key)
}

func (env *Environment) GetIntSlice(key string) []int {
	return env.viper.GetIntSlice(key)
}

func (env *Environment) GetFloat64(key string) float64 {
	return env.viper.GetFloat64(key)
}

func (env *Environment) GetTime(key string) time.Time {
	return env.viper.GetTime(key)
}

func (env *Environment) GetDuration(key string) time.Duration {
	return env.viper.GetDuration(key)
}

func (env *Environment) Viper() *v.Viper {
	return env.viper
}

func loadDotEnv(path string) {
	if path != "" {
		if _, err := os.Stat(path); err == nil {
			err := godotenv.Load(path)
			if err != nil {
				panic(err)
			}
		}
	}
}

func createViper(prefix string) *v.Viper {
	env := v.New()
	env.SetEnvPrefix(prefix)
	env.AutomaticEnv()
	return env
}
