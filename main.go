package main

import (
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
	runDBMigration(extConfig.MigrationURL, extConfig.DBSource)
	runGinServer(extConfig)
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}

func runGinServer(extConfig util.ExtConfig) {
	server, err := api.NewServer(extConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(extConfig.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}
