package main

import (
	"github.com/gCastl/go-clean-template/internal/config"
	"github.com/gCastl/go-clean-template/internal/interface/api"
	"github.com/gCastl/go-clean-template/pkg/clicfg"
)

func main() {
	cli := clicfg.New(
		clicfg.String("port", "PORT", "8080"),
		clicfg.String("log-level", "LOG_LEVEL", "info"),
	)

	conf, err := config.NewConfig(
		config.WithPort(cli.GetString("port")),
		config.WithLogLevel(cli.GetString("log-level")),
	)
	if err != nil {
		panic(err)
	}

	if err := api.NewAPI(conf).ListenAndServe(); err != nil {
		panic(err)
	}
}
