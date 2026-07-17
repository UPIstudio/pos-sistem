package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/luthfi/pos/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	config.ConnectDB()
	app := fiber.New()

	log.Println("Server berjalan di :8080")
	log.Fatal(app.Listen(":8080"))

	select {}
}
