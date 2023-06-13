package config

import (
	"fmt"
	"os"

	"githiub.com/qodirtok/go-rest-login/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect()  {
	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, username, password, databaseName, port)

	models.Database , err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
	panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
		Migration()
	}
}

func Migration()  {
	models.Database.AutoMigrate(&models.User{})
}