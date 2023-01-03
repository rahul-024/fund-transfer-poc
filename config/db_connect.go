package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDatabase(appConfig *AppConfig) {
	dsn := appConfig.Datasource.Dsn
	switch appConfig.Datasource.DbType {
	case "postgres":
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "mysql":
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}
	if err != nil {
		panic("Failed to connect to database")
	}
	if err != nil {
		return
	}
}
