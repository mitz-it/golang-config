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

func NewConfig(cfg StartConfig) *Config {
	if cfg.ConfigPath != "" {
		if _, err := os.Stat(cfg.ConfigPath); err == nil {
			err := godotenv.Load(cfg.ConfigPath)
			if err != nil {
				panic(err)
			}
		}
	}
	config := viper.New()
	config.SetEnvPrefix(cfg.Prefix)
	config.AutomaticEnv()

	return &Config{Standard: config}
}
