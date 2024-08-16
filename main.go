package main

import (
	"log"

	"github.com/hazaloolu/GoRestApi/database"
	"github.com/hazaloolu/GoRestApi/model"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
}

func loadDatabase() {
	var err error
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})

	if err != nil {
		panic(err)
	}

}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
