package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	_, err = database.DB()
	if err != nil {
		log.Fatal("Error closing database: ", err)
	}

	/*defer sqlDB.Close()*/

	err = database.AutoMigrate(&Dream{}, &DreamImageQueue{})
	if err != nil {
		log.Fatal("Error migrating database: ", err)
	}

	DB = database
}
