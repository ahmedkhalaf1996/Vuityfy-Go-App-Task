package models

import (
	"log"

	"gorm.io/gorm/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	var err error
	var database *gorm.DB

	// Connect to the SQLite database
	database, err = gorm.Open(sqlite.Open("golang.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate
	// database.AutoMigrate(&Posts{})
	database.AutoMigrate(&Categories{})
	database.AutoMigrate(&SubCategory{})
	database.AutoMigrate(&Item{})
	database.AutoMigrate(&Order{})
	database.AutoMigrate(OrderItem{})

	DB = database
}
