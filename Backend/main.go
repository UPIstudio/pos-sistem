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

	memRepo := repository.NewMemberRepository()
	memSvc := service.NewMemberService(memRepo)
	memHdl := handler.NewMemberHandler(memSvc)

	mem := app.Group("/api/member", middleware.AuthMiddleware)
	mem.Get("/", memHdl.GetAll)
	mem.Get("/search", memHdl.Search)
	mem.Post("/create", memHdl.Create)
	mem.Put("/update/:id", memHdl.Update)
	mem.Delete("/delete/:id", memHdl.Delete)

	transRepo := repository.NewTransactionRepository()
	transSvc := service.NewTransactionService(transRepo)
	transHdl := handler.NewTransactionHandler(transSvc)

	trans := app.Group("/api/transaksi", middleware.AuthMiddleware)

	trans.Get("/", transHdl.GetAll)
	trans.Get("/detail/:id", transHdl.GetByID)
	trans.Post("/create", transHdl.Create)
	trans.Put("/status/:id", transHdl.UpdateStatus)
	trans.Delete("/delete/:id", transHdl.Delete)

	reportSvc := service.NewReportService(transRepo)
	reportHdl := handler.NewReportHandler(reportSvc)

	report := app.Group("/api/laporan", middleware.AuthMiddleware)
	report.Get("/export", reportHdl.ExportExcel)

	log.Println("Server berjalan di :8080")
	app.Listen(":8080")

	select {}
}
