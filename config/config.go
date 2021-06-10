package config

import (
	"github.com/gofiber/fiber/v2"

	"github.com/confetti-framework/support/env"
)

var Fiber *fiber.Config

var App = struct {
	Name string
	Url  string
	Port int
}{
	Name: env.StringOr("APP_NAME", "Fiber"),
	Url:  env.StringOr("APP_URL", "http://localhost"),
	Port: env.IntOr("APP_PORT", 3000),
}

var Cors = struct {
	Origins string
}{
	Origins: env.StringOr("CORS_ORIGINS", "*"),
}

var Database = struct {
	Url      string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
}{
	Url:      env.String("DATABASE_URL"),
	Host:     env.StringOr("DB_HOST", "127.0.0.1"),
	Port:     env.IntOr("DB_PORT", 3306),
	Name:     env.StringOr("DB_DATABASE", "fiber"),
	Username: env.StringOr("DB_USERNAME", "fiber"),
	Password: env.StringOr("DB_PASSWORD", ""),
}
