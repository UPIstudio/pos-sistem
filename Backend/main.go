package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/luthfi/pos/config"
	"github.com/luthfi/pos/handler"
	"github.com/luthfi/pos/middleware"
)

func main() {
	config.ConnectDB()
	app := fiber.New()
	godotenv.Load()

	app.Post("/register", handler.RegisterHandler)
	app.Post("/login", handler.LoginHandler)

	api := app.Group("/api", middleware.AuthMiddleware)
	api.Get("/me", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"user": c.Locals("username")})
	})

	log.Println("Server berjalan di :8080")
	app.Listen(":8080")

	select {}
}
