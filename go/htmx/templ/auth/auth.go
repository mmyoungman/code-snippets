package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"mmyoungman/templ/utils"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
	EndSessionURL string
}

func Setup() (*Authenticator, error) {
	provider := loadProvider()

	// To get end_session_endpoint. See https://github.com/coreos/go-oidc/pull/226#issuecomment-1130411016
	var claims struct {
		EndSessionURL string `json:"end_session_endpoint"`
	}
	err := provider.Claims(&claims)
	if err != nil {
		log.Println("Didn't find end_session_endpoint in discovery?")
		return nil, err
	}

	callbackURL := fmt.Sprintf("%s%s", utils.GetPublicURL(), "/auth/callback")

	conf := oauth2.Config{
		ClientID:     utils.Getenv("OIDC_CLIENT_ID"),
		ClientSecret: utils.Getenv("OIDC_CLIENT_SECRET"),
		RedirectURL:  callbackURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	authObj := Authenticator{
		Provider:      provider,
		Config:        conf,
		EndSessionURL: claims.EndSessionURL,
	}

	return &authObj, nil
}

func loadProvider() *oidc.Provider {
	numRetries := 10
	retryWaitSecs := 15

	provider, err := oidc.NewProvider(
		context.Background(),
		// @MarkFix create Config object to store this kind of thing so it doesn't have to be constructed all over the place
		utils.Getenv("OIDC_URL"),
	)

	// retries
	for i := 0; i < numRetries && provider == nil; i++ {
		if err != nil {
			slog.Error("Failed to load auth provider. Retrying after waiting...", "retryWaitSecs", retryWaitSecs)
			time.Sleep(time.Duration(retryWaitSecs) * time.Second)
		}

		provider, err = oidc.NewProvider(
			context.Background(),
			// @MarkFix create Config object to store this kind of thing so it doesn't have to be constructed all over the place
			utils.Getenv("OIDC_URL"),
		)
	}

	if provider == nil {
		log.Fatal("Failed to load provider after ", numRetries, " retries")
	}

	return provider

}

// VerifyIDToken verifies that an *oauth2.Token is a valid *oidc.IDToken.
func (a *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

func GenerateRandomState() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		slog.Error("Failed to generate random state", "error", err)
	}

	return base64.StdEncoding.EncodeToString(b)
}
