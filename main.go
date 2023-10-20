package main

import (
	"github.com/joho/godotenv"
	"github.com/yazilimcigenclik/dream-ai-backend/app/router"
	"github.com/yazilimcigenclik/dream-ai-backend/config"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	config.InitLog()
}

func main() {

	port := os.Getenv("APP_PORT")

	log.Println("Starting server on port " + port)

	db := config.ConnectToDB()

	init := config.NewInitialization(db)

	log.Println("Initializing...")

	app := router.Init(init)

	err := app.Run(":" + port)
	if err != nil {
		log.Fatal("Error running server")
		return
	}

}
