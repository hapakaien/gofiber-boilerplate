package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/hapakaien/fiber-boilerplate/config"
	"github.com/hapakaien/fiber-boilerplate/handlers"
)

func main() {
	// Use an external setup function in order
	// to configure the app in tests as well
	app := Setup()

	// Start server
	log.Fatal(app.Listen(":" + config.Config("PORT")))
}

// Setup func to config Fiber app
func Setup() *fiber.App {
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
		c.SendStatus(404)
		return c.JSON(fiber.Map{
			"success": false,
			"message": "404 not found",
			"data":    "",
		})
	})

	// Return the configured app
	return app
}
