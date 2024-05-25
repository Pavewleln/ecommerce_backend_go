package main

import (
	"github.com/joho/godotenv"
	"log"
	"main/src/config"
	"main/src/routes"
	"main/src/utils"
	"os"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.LoadEnv()
	config.ConnectDB()
	router := routes.SetupRoutes()

	router.Run(os.Getenv("PORT"))
}
