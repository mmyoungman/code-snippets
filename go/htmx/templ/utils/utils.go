package utils

import (
	"fmt"
	"log"
	"os"
)

type reqCtxKey int

const (
	ReqUserCtxKey reqCtxKey = iota
	CspNonceCtxKey reqCtxKey = iota
)

func Getenv(key string) string {
	variable := os.Getenv(key)

	if variable == "" {
		log.Fatalf("Failed to get environment variable \"%s\". Is it added to .env?", key)
	}

	return variable
}

func GetPublicURL() string { // @hotreload
	if !IsProd && os.Getenv("PROXY_URL") != "" { // ensure !IsProd to prevent shenanigans
		return os.Getenv("PROXY_URL")
	}

	return fmt.Sprintf("%s:%s", Getenv("PUBLIC_HOST"), Getenv("PUBLIC_PORT"))
}
