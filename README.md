# Golang - Config

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=mitz-it_golang-config&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=mitz-it_golang-config) [![Coverage](https://sonarcloud.io/api/project_badges/measure?project=mitz-it_golang-config&metric=coverage)](https://sonarcloud.io/summary/new_code?id=mitz-it_golang-config)

Abstraction over [Viper](https://github.com/spf13/viper) to read environment variables.

Environment variables from `.env` files are also supported.

## Installation

```bash
  go get -u github.com/mitz-it/golang-config
```

## Usage 1

```go
package main

import "github.com/mitz-it/golang-config"

func main() {
  prefix := "mtz" // all environment variables starting with MTZ_ will be loaded
  start := config.StartConfig{
    ConfigPath: "../config/.env",
    Prefix: prefix,
  }

  cfg := config.NewConfig(start)

  connStr := cfg.Standard.GetString("connection_string") // MTZ_CONNECTION_STRING

  // ...
}
```

## Usage 2

```go
package main

import "github.com/mitz-it/golang-config"

func init() {
  prefix := "mtz" // all environment variables starting with MTZ_ will be loaded
  path := "../config/.env"
  config.LoadEnv(prefix, path)
}

func main() {
  connStr := config.Env.GetString("connection_string") // MTZ_CONNECTION_STRING

  // ...
}
```

## Scopes

```go
package main

import "github.com/mitz-it/golang-config"

func init() {
  config.LoadScopedEnv("scope1", "sc1", "../config/.env")
  config.LoadScopedEnv("scope1", "sc2", "../config/.env")
}

func main() {
  connStr1 := config.Scope("scope1").Env.GetString("connection_string") // SC1_CONNECTION_STRING
  connStr2 := config.Scope("scope2").Env.GetString("connection_string") // SC2_CONNECTION_STRING

  // ...
}
```
