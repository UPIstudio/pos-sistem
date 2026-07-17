package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/luthfi/pos/repository"
)

func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if len(tokenString) < 7 {
		return c.Status(401).JSON(fiber.Map{"error": "Harus Login Terlebih Dahulu"})
	}

	tokenString = tokenString[7:]

	token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		username := claims["username"].(string)
		user, _ := repository.FindByUsername(username)

		if user.Token != tokenString {
			return c.Status(401).JSON(fiber.Map{"error": "Token Tidak Valid"})
		}

		c.Locals("username", username)
		return c.Next()
	}

	return c.Status(401).JSON(fiber.Map{"error": "Token Tidak Valid"})

}
