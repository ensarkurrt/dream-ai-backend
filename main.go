package main

import (
	"github.com/joho/godotenv"
	"github.com/yazilimcigenclik/dream-ai-backend/controllers"
	"github.com/yazilimcigenclik/dream-ai-backend/models"
	"log"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	models.ConnectDatabase()

	controllers.New()

	if err != nil {
		log.Fatal("error with ListenAndServe on main", err)
	}
}
