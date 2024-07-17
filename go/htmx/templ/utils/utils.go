package utils

import (
	"fmt"
	"log"
	"os"
)

type reqCtxKey int

const (
	ReqUserCtxKey reqCtxKey = iota
)

func Getenv(key string) string {
	variable := os.Getenv(key)

	if variable == "" {
		log.Fatalf("Failed to get environment variable \"%s\". Is it added to .env?", key)
	}

	return variable
}

func GetPublicURL() string { // @hotreload
	// @MarkFix rename TEMPL_WATCH_PROXY_URL to just PROXY_URL?
	if !IsProd && os.Getenv("TEMPL_WATCH_PROXY_URL") != "" { // ensure !IsProd to prevent shenanigans
		return os.Getenv("TEMPL_WATCH_PROXY_URL")
	}

	return fmt.Sprintf("%s:%s", Getenv("PUBLIC_HOST"), Getenv("PUBLIC_PORT"))
}
