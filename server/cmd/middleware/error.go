package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.Handler {

	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
			return nil
		}
		return nil
	}
}
