package Utils

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Failed to load .env file!")
	}
}

func GetPort() string {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8888"
	}
	return port
}
