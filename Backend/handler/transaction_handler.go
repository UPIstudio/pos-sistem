package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/luthfi/pos/models"
	"github.com/luthfi/pos/service"
)

type TransactionHandler struct{ svc *service.TransactionService }

func NewTransactionHandler(s *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{svc: s}
}

func (h *TransactionHandler) Create(c *fiber.Ctx) error {
	var t models.Transaction
	if err := c.BodyParser(&t); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}
	if err := h.svc.Create(&t); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"message": "Transaksi berhasil"})
}

func (h *TransactionHandler) GetAll(c *fiber.Ctx) error {
	data, _ := h.svc.GetAll()
	return c.JSON(data)
}

func (h *TransactionHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	data, err := h.svc.GetByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Tidak ditemukan"})
	}
	return c.JSON(data)
}

func (h *TransactionHandler) UpdateStatus(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	type Request struct {
		Status string `json:"status"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.svc.UpdateStatus(uint(id), req.Status); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Status diupdate ke " + req.Status})
}

func (h *TransactionHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	h.svc.Delete(uint(id))
	return c.JSON(fiber.Map{"message": "Transaksi dihapus"})
}
