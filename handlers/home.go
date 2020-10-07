package handlers

import "github.com/gofiber/fiber/v2"

// Home func to root endpoint
func Home(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Yo!",
		"data": "",
	})
}