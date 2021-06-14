package routes

import (
	"github.com/hapakaien/fiber-boilerplate/controllers"
	"github.com/hapakaien/fiber-boilerplate/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App) {
	// Middleware
	api := app.Group("/", middlewares.Cors(), middlewares.Compress(), func(c *fiber.Ctx) error {
		c.Accepts("application/json")

		return c.Next()
	})
	api.Get("/", controllers.Home)
}
