package middleware

import (
	"refresh/internal/app/config"
	"refresh/pkg/auth"

	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")

		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		a := auth.Token{}

		id, err := a.ParseToken(tokenString, config.InitConfig().ACCES_TOKEN)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
		}

		c.Locals("user", fiber.Map{
			"id": id,
		})

		return c.Next()
	}
}
