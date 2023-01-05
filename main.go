package main

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"github.com/rahul-024/fund-transfer-poc/config"
	logFactory "github.com/rahul-024/fund-transfer-poc/loggerfactory"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/gorm"
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
	err = viper.Unmarshal(&config.AppConf)
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
		err = viper.Unmarshal(&config.AppConf)
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
	db := config.ConnectDatabase(&config.AppConf)
	runDBMigration(&config.AppConf)
	loadLogger(config.AppConf.Log)
	runGinServer(&config.AppConf, db)
}

func runDBMigration(appConfig *config.AppConfig) {
	migration, err := migrate.New(appConfig.DbMigrationPath, appConfig.Datasource.Dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}

func runGinServer(appConfig *config.AppConfig, db *gorm.DB) {
	server, err := config.NewServer(db)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(appConfig.ServerConfig.HttpServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}

// loads the logger
func loadLogger(lc config.LogConfig) error {
	loggerType := lc.Code
	err := logFactory.GetLogFactoryBuilder(loggerType).Build(&lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
