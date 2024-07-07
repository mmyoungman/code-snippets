package auth

import (
	"context"
	"errors"
	"mmyoungman/templ/utils"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

var State string
var AccessToken string
var RawIDToken string
var Profile map[string]interface{}

func Setup() (*Authenticator, error) {
	provider, err := oidc.NewProvider( // @MarkFix replace this with custom so I get things like end_session_endpoint from discovery url response?
		context.Background(),
		utils.Getenv("KEYCLOAK_URL") + "/realms/" + utils.Getenv("KEYCLOAK_REALM"),
	)
	if err != nil {
		return nil, err
	}

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