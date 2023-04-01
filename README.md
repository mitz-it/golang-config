# Golang - Config

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
  prefix := "MTZ" // all environment variables starting with MITZ_ will be loaded
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
  prefix := "MTZ" // all environment variables starting with MITZ_ will be loaded
  path := "../config/.env"
  config.Load(prefix, path)
}

func main() {
  connStr := config.Instance.GetString("connection_string") // MTZ_CONNECTION_STRING

  // ...
}
```
