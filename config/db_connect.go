package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error

func ConnectDatabase(appConfig *AppConfig) (db *gorm.DB) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // enable color
		},
	)
	dsn := appConfig.Datasource.Dsn
	switch appConfig.Datasource.DbType {
	case "postgres":
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})

	case "mysql":
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}
	if err != nil {
		panic("Failed to connect to database")
	}
	if err != nil {
		return
	}
	return DB
}
