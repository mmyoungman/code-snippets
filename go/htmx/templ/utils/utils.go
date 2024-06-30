package utils

import (
	"log"
	"os"
)

func Getenv(key string) string {
	variable := os.Getenv(key)

	if variable == "" {
		log.Fatal("Failed to get environment variable \"" + key + "\". Is it added to .env?")
	}

	return variable
}