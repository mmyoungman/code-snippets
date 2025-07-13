package utils

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"mmyoungman/templ/database/jet/model"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/google/uuid"
)

type reqCtxKey int

const (
	UserCtxKey      reqCtxKey = iota
	CspNonceCtxKey  reqCtxKey = iota
	CsrfTokenCtxKey reqCtxKey = iota
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

func GetContextCSRFToken(r *http.Request) string {
	csrfUntyped := r.Context().Value(CsrfTokenCtxKey)
	if csrfUntyped == nil {
		log.Fatal("CSRF token didn't get set")
	}
	return csrfUntyped.(string)
}

func ValidateCSRFToken(r *http.Request, csrfToken string) bool { // @MarkFix use this! And test it
	expectedCsrfToken := GetContextCSRFToken(r)

	// refresh token
	newToken := uuid.New().String()
	// @MarkFix IMPORTANT save the csrfToken to the secure cookie!
	SetContextValue(r, CsrfTokenCtxKey, newToken)

	return csrfToken == expectedCsrfToken
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

func GenerateRandomStr() string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	newNonce := make([]rune, 20)
	for i := 0; i < 20; i++ {
		newNonce[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(newNonce)
}
