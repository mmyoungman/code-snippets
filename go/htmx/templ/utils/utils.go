package utils

import (
	"fmt"
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

func GetPublicURL() string { // @hotreload
	if !IsProd && os.Getenv("TEMPL_WATCH_PROXY_URL") != "" { // ensure !IsProd to prevent shenanigans
		return os.Getenv("TEMPL_WATCH_PROXY_URL")
	}

	return fmt.Sprintf("%s:%s", os.Getenv("PUBLIC_HOST"), os.Getenv("PUBLIC_PORT"))
}