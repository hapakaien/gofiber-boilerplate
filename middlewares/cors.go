package middlewares

import (
	"github.com/hapakaien/fiber-boilerplate/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors() fiber.Handler {
  	return cors.New(cors.Config{
  		AllowOrigins:     config.Cors.Origins,
  		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
  		AllowHeaders:     "",
  		AllowCredentials: false,
  		ExposeHeaders:    "",
  		MaxAge:           0,
  	})
}
