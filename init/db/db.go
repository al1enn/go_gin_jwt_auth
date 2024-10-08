package db

import (
	"auth/pkg/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(url string) {
	var err error
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database.")
	}
	// Автоматическая миграция моделей
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Panic(err)
	}

}
