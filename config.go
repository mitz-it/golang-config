package config

import (
	v "github.com/spf13/viper"
)

type StartConfig struct {
	Prefix     string
	ConfigPath string
}

type Config struct {
	Standard *v.Viper
}

func NewConfig(cfg StartConfig) *Config {
	loadDotEnv(cfg.ConfigPath)

	config := createViper(cfg.Prefix)

	return &Config{Standard: config}
}
