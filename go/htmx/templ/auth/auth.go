package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"mmyoungman/templ/utils"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

// @MarkFix replace with session state
var State string
var AccessToken string
var RefreshToken string
var TokenType string
var RawIDToken string
var Expiry time.Time
var Profile map[string]interface{}
var EndSessionURL string

func Setup() (*Authenticator, error) {
	provider, err := oidc.NewProvider(
		context.Background(),
		// @MarkFix create Config object to store this kind of thing so it doesn't have to be constructed all over the place
		utils.Getenv("KEYCLOAK_URL") + "/realms/" + utils.Getenv("KEYCLOAK_REALM"),
	)
	if err != nil {
		return nil, err
	}

	// Hack to get end_session_endpoint https://github.com/coreos/go-oidc/pull/226#issuecomment-1130411016
	var claims struct {
		EndSessionURL string `json:"end_session_endpoint"`
	}
	err = provider.Claims(&claims)
	if err != nil {
		log.Println("Didn't find end_session_endpoint in discovery?")
		return nil, err
	}
	EndSessionURL = claims.EndSessionURL

	conf := oauth2.Config{
		ClientID: utils.Getenv("KEYCLOAK_CLIENT_ID"),
		ClientSecret: utils.Getenv("KEYCLOAK_CLIENT_SECRET"),
		RedirectURL: utils.Getenv("KEYCLOAK_CALLBACK_URL"),
		Endpoint: provider.Endpoint(),
		Scopes: []string{oidc.ScopeOpenID, "profile"},
	}

	authObj := Authenticator{
		Provider: provider,
		Config: conf,
	}

	return &authObj, nil
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

func GenerateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}