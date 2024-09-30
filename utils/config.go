package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	env := os.Getenv("APP_ENV")
	envFile := fmt.Sprintf("./config/.env.%s", env)

	err := godotenv.Load(envFile)

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func GetAppEnv() string {
	return os.Getenv("APP_ENV")
}
