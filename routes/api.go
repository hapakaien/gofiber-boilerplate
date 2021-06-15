package routes

import (
	"github.com/hapakaien/fiber-boilerplate/controllers"
	"github.com/hapakaien/fiber-boilerplate/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App) {
	// Middleware
	api := app.Group("/", func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		return c.Next()
	})
	api.Use(middlewares.Cors())
	api.Use(middlewares.Favicon())
	api.Use(middlewares.Compress())
	api.Use(middlewares.Etag())

	// Home
	api.Get("/", controllers.Home)
}
