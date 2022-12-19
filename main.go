package main

import (
	"os"

	"github.com/rahul-024/fund-transfer-poc/api"
	"github.com/rahul-024/fund-transfer-poc/config"
	"github.com/rahul-024/fund-transfer-poc/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	extConfig, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	if extConfig.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	config.ConnectDatabase(extConfig.DBSource)
}

func runGinServer(config util.Config) {
	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}
