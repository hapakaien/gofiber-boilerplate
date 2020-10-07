package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/compress"

	"github.com/hapakaien/gofiber-boilerplate/config"
)

func main() {
	// Fiber instance
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: config.Config("CORS_DOMAIN"),
	}))

	// Comporess
	app.Use(compress.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	// Start server
	app.Listen(":3000")
}