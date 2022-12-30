package main

import (
	"github.com/TulioGuaraldoB/cross-roads/config/env"
	"github.com/TulioGuaraldoB/cross-roads/infra/server"
)

func main() {
	env.GetEnvironmentVariables()

	server := server.New()
	server.Run()
}
