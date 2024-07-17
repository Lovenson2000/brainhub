package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func Compresssor() fiber.Handler {
	return compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	})
}
