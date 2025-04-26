package env

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// TODO: TUNE THIS
func Init() error {
	if os.Getenv("APP_ENV") == "production" {
		return nil
	}

	err := godotenv.Load(".env")
	if err != nil {
		return errors.New("Error loading .env file")
	}
	return nil
}

func Get(envName string) string {
	envVariable := os.Getenv(envName)
	if envVariable == "" {
		log.Panicf("The .env variable: %s wasn't found", envName)
	}

	return envVariable
}
