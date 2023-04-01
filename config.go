package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type StartConfig struct {
	Prefix     string
	ConfigPath string
}

type Config struct {
	Standard *viper.Viper
}

var Instance *viper.Viper

func NewConfig(cfg StartConfig) *Config {
	loadEnv(cfg.ConfigPath)

	config := configureViper(cfg.Prefix)

	return &Config{Standard: config}
}

func Load(prefix, path string) *viper.Viper {
	loadEnv(path)

	Instance = configureViper(prefix)

	return Instance
}

func loadEnv(path string) {
	if path != "" {
		if _, err := os.Stat(path); err == nil {
			err := godotenv.Load(path)
			if err != nil {
				panic(err)
			}
		}
	}
}

func configureViper(prefix string) *viper.Viper {
	config := viper.New()
	config.SetEnvPrefix(prefix)
	config.AutomaticEnv()
	return config
}
