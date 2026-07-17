package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/luthfi/pos/models"
	"github.com/luthfi/pos/service"
)

type ProductHandler struct {
	svc *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{svc: s}
}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	var p models.Product
	c.BodyParser(&p)
	if err := h.svc.CreateProduct(&p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Gagal"})
	}
	return c.JSON(fiber.Map{"message": "Produk dibuat"})
}

func (h *ProductHandler) GetAll(c *fiber.Ctx) error {
	data, _ := h.svc.GetAll()
	return c.JSON(data)
}

func (h *ProductHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var p models.Product
	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON tidak valid"})
	}

	if err := h.svc.UpdateProduct(id, &p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Gagal update"})
	}
	return c.JSON(fiber.Map{"message": "Produk berhasil diupdate"})
}

func (h *ProductHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	h.svc.Delete(id)
	return c.JSON(fiber.Map{"message": "Dihapus"})
}

func (h *ProductHandler) Search(c *fiber.Ctx) error {
	query := c.Query("q")
	data, _ := h.svc.Search(query)
	return c.JSON(data)
}
