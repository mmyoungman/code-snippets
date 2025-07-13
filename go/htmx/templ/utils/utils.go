package utils

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"mmyoungman/templ/database/sqlc_gen"
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

type BaseArgs struct {
	Username  string
	Nonce     string
	CsrfToken string
}

func SetContextValue(r *http.Request, key reqCtxKey, value any) {
	ctx := r.Context()
	newCtx := context.WithValue(ctx, key, value)
	*r = *r.WithContext(newCtx)
}

func GetContextUser(r *http.Request) *database.User {
	userUntyped := r.Context().Value(UserCtxKey)
	if userUntyped == nil {
		return nil
	}

	if user, ok := userUntyped.(*database.User); ok {
		return user
	}

	// @MarkFix if there is a user in session but not in db, we reach here. Maybe should validate the user rather than trusting session?
	return nil
}

func GetContextCspNonce(r *http.Request) string {
	nonceUntyped := r.Context().Value(CspNonceCtxKey)
	if nonceUntyped == nil {
		log.Fatal("CSP nonce didn't get set")
	}
	return nonceUntyped.(string)
}

func GenerateBaseArgs(r *http.Request) BaseArgs {
	firstName := ""

	user := GetContextUser(r)
	if user != nil {
		firstName = user.Firstname
	}

	return BaseArgs{
		Nonce:     GetContextCspNonce(r),
		CsrfToken: GetContextCSRFToken(r),
		Username:  firstName,
	}
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
