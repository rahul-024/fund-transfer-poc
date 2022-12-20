package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(dbSource string) {
	database, err := gorm.Open(postgres.Open(dbSource), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	if err != nil {
		return
	}
	DB = database
}
