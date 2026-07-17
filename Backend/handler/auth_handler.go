package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/luthfi/pos/service"
)

func RegisterHandler(c *fiber.Ctx) error {
	type Req struct{ Username, Password, Role string }
	var r Req
	c.BodyParser(&r)

	err := service.Register(r.Username, r.Password, r.Role)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Register Gagal"})
	}

	return c.JSON(fiber.Map{"message": "Register Berhasil | Silahkan Login"})
}

func LoginHandler(c *fiber.Ctx) error {
	type Req struct{ Username, Password, Role string }
	var r Req
	c.BodyParser(&r)

	fmt.Printf("DEBUG: Request Login untuk username: %s\n", r.Username)

	if r.Username == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Username tidak boleh kosong"})
	}

	token, err := service.Login(r.Username, r.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Login Gagal"})
	}

	return c.JSON(fiber.Map{"token": token})
}
