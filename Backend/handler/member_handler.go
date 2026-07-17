package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/luthfi/pos/models"
	"github.com/luthfi/pos/service"
)

type MemberHandler struct {
	svc *service.MemberService
}

func NewMemberHandler(s *service.MemberService) *MemberHandler {
	return &MemberHandler{svc: s}
}

func (h *MemberHandler) GetAll(c *fiber.Ctx) error {
	data, _ := h.svc.GetAll()
	return c.JSON(data)
}

func (h *MemberHandler) Search(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Parameter pencarian 'q' tidak boleh kosong"})
	}

	data, err := h.svc.Search(query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mencari member"})
	}
	return c.JSON(data)
}

func (h *MemberHandler) Create(c *fiber.Ctx) error {
	var m models.Member
	if err := c.BodyParser(&m); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format JSON tidak valid"})
	}

	if err := h.svc.Create(&m); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat member, pastikan Nomor HP belum terdaftar"})
	}
	return c.Status(201).JSON(fiber.Map{"message": "Member berhasil dibuat", "data": m})
}

func (h *MemberHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var m models.Member
	if err := c.BodyParser(&m); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format JSON tidak valid"})
	}

	if err := h.svc.Update(uint(id), &m); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengupdate member"})
	}
	return c.JSON(fiber.Map{"message": "Member berhasil diupdate"})
}

func (h *MemberHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.svc.Delete(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus member"})
	}
	return c.JSON(fiber.Map{"message": "Member berhasil dihapus"})
}
