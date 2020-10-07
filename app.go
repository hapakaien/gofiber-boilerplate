package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber instance
	app := fiber.New()

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