package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/compress"

	"github.com/hapakaien/gofiber-boilerplate/config"
	"github.com/hapakaien/gofiber-boilerplate/handlers"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.Config("CORS_DOMAIN"),
	}))
	app.Use(compress.New())

	// Routes
	api := app.Group("/", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.Next()
	})
	api.Get("/", handlers.Home)

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	// Start server
	app.Listen(":3000")
}