package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		stop := time.Now()

		log.Printf("[%s] %s %s %d %s",
			c.IP(),
			c.Method(),
			c.OriginalURL(),
			c.Response().StatusCode(),
			stop.Sub(start).String(),
		)

		return err
	}
}
