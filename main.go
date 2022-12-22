package main

import (
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rahul-024/fund-transfer-poc/config"
	db "github.com/rahul-024/fund-transfer-poc/db/config"
	"github.com/rahul-024/fund-transfer-poc/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//	@title			Fund transfer service
//	@version		1.0
//	@description	A rest based service in Go using Gin framework.
//	@termsOfService	https://tos.iexceed.dev

//	@contact.name	Iexceed technology solutions
//	@contact.url	https://www.i-exceed.com/contact-us/
//	@contact.email	rahul.r@i-exceed.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8081
//	@BasePath	/api/v1

func main() {
	extConfig, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	if extConfig.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	db.ConnectDatabase(extConfig.DBSource)
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
	server, err := config.NewServer(extConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(extConfig.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}
