package main

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rahul-024/fund-transfer-poc/config"
	db "github.com/rahul-024/fund-transfer-poc/db/config"
	"github.com/rahul-024/fund-transfer-poc/models"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {
	profile := initProfile()
	LoadConfig(profile)
}

func initProfile() string {
	var profile string
	profile = os.Getenv("ENV_PROFILE")
	if len(profile) <= 0 {
		profile = "local"
	}
	fmt.Println("ENV_PROFILE: " + profile)
	return profile
}

func LoadConfig(profile string) {
	viper.AddConfigPath("./profiles")
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&models.RuntimeConf)
	if err != nil {
		panic(err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		var err error
		err = viper.ReadInConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = viper.Unmarshal(&models.RuntimeConf)
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	viper.WatchConfig()
}

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
	db.ConnectDatabase(&models.RuntimeConf)
	runDBMigration(&models.RuntimeConf)
	runGinServer(&models.RuntimeConf)
}

func runDBMigration(runtimeConfig *models.RuntimeConfig) {
	migration, err := migrate.New(runtimeConfig.DbMigrationPath, runtimeConfig.Datasource.Dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}

func runGinServer(runtimeConfig *models.RuntimeConfig) {
	server, err := config.NewServer()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(runtimeConfig.Server.HttpServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}
