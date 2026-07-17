package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/luthfi/pos/models"
	"github.com/luthfi/pos/service"
)

type CategoryHandler struct{ svc *service.CategoryService }

func NewCategoryHandler(s *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{svc: s}
}

func (h *CategoryHandler) GetAll(c *fiber.Ctx) error {
	data, _ := h.svc.GetAll()
	return c.JSON(data)
}

func (h *CategoryHandler) Create(c *fiber.Ctx) error {
	var cat models.Category
	if err := c.BodyParser(&cat); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}
	h.svc.Create(&cat)
	return c.JSON(fiber.Map{"message": "Kategori dibuat"})
}

func (h *CategoryHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var cat models.Category
	c.BodyParser(&cat)
	h.svc.Update(uint(id), &cat)
	return c.JSON(fiber.Map{"message": "Kategori diupdate"})
}

func (h *CategoryHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	h.svc.Delete(uint(id))
	return c.JSON(fiber.Map{"message": "Kategori dihapus"})
}
