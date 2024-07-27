package utils

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
)

type reqCtxKey int

const (
	UserCtxKey     reqCtxKey = iota
	CspNonceCtxKey reqCtxKey = iota
)

func GetContextCspNonce(r *http.Request) string {
	nonceUntyped := r.Context().Value(CspNonceCtxKey)
	if nonceUntyped == nil {
		log.Fatal("CSP Nonce didn't get set")
	}
	return nonceUntyped.(string)
}

func SetContextValue(r *http.Request, key reqCtxKey, value any) {
	ctx := r.Context()
	newCtx := context.WithValue(ctx, key, value)
	*r = *r.WithContext(newCtx)
}

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
