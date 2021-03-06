package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
)

func Etag() fiber.Handler {
	return etag.New(etag.Config{
		Weak: false,
	})
}
