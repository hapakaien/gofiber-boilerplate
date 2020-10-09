package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("Error loading .env file")
	}
	return os.Getenv(key)
}
