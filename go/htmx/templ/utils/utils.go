package utils

import (
	"context"
	"fmt"
	"log"
	"mmyoungman/templ/database/jet/model"
	"net/http"
	"os"
	"runtime/debug"
)

type reqCtxKey int

const (
	UserCtxKey     reqCtxKey = iota
	CspNonceCtxKey reqCtxKey = iota
)

func SetContextValue(r *http.Request, key reqCtxKey, value any) {
	ctx := r.Context()
	newCtx := context.WithValue(ctx, key, value)
	*r = *r.WithContext(newCtx)
}

func GetContextUser(r *http.Request) *model.User {
	userUntyped := r.Context().Value(UserCtxKey)
	if userUntyped == nil {
		return nil
	}
	return userUntyped.(*model.User)
}

func GetContextCspNonce(r *http.Request) string {
	nonceUntyped := r.Context().Value(CspNonceCtxKey)
	if nonceUntyped == nil {
		log.Fatal("CSP nonce didn't get set")
	}
	return nonceUntyped.(string)
}

func Getenv(key string) string {
	variable := os.Getenv(key)

	if variable == "" {
		log.Fatalf("Failed to get environment variable \"%s\". Is it added to .env?\n%s", key, debug.Stack())
	}

	return variable
}

func GetPublicURL() string { // @hotreload
	if !IsProd && os.Getenv("PROXY_URL") != "" { // ensure !IsProd to prevent shenanigans
		return os.Getenv("PROXY_URL")
	}

	return fmt.Sprintf("%s:%s", Getenv("PUBLIC_HOST"), Getenv("PUBLIC_PORT"))
}
