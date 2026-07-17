package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/luthfi/pos/config"
	"github.com/luthfi/pos/handler"
	"github.com/luthfi/pos/middleware"
	"github.com/luthfi/pos/repository"
	"github.com/luthfi/pos/service"
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

	prodRepo := repository.NewProductRepository()
	prodSvc := service.NewProductService(prodRepo)
	prodHdl := handler.NewProductHandler(prodSvc)

	prod := app.Group("/api/produk", middleware.AuthMiddleware)
	prod.Get("/", prodHdl.GetAll)
	prod.Get("/search", prodHdl.Search)
	prod.Post("/create", prodHdl.Create)
	prod.Put("/update/:id", prodHdl.Update)
	prod.Delete("/delete/:id", prodHdl.Delete)

	catRepo := repository.NewCategoryRepository()
	catSvc := service.NewCategoryService(catRepo)
	catHdl := handler.NewCategoryHandler(catSvc)

	cat := app.Group("/api/kategori", middleware.AuthMiddleware)
	cat.Get("/", catHdl.GetAll)
	cat.Post("/create", catHdl.Create)
	cat.Put("/edit/:id", catHdl.Update)
	cat.Delete("/delete/:id", catHdl.Delete)

	log.Println("Server berjalan di :8080")
	app.Listen(":8080")

	select {}
}
