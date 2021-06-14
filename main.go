package main

import (
	"fmt"
	"log"

	"github.com/hapakaien/fiber-boilerplate/config"
	"github.com/hapakaien/fiber-boilerplate/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Use an external setup function in order
	// to configure the app in tests as well
	app := setup()

	// Start server
	log.Fatal(app.Listen(fmt.Sprint(":", config.App.Port)))
}

// Setup func to config Fiber app
func setup() *fiber.App {
	// Fiber instance
	app := fiber.New()

	// Route
	routes.Api(app)

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		c.SendStatus(fiber.StatusNotFound)

		return c.JSON(fiber.Map{
			"success": false,
			"message": "404 not found",
			"data":    "",
		})
	})

	// Return the configured app
	return app
}
