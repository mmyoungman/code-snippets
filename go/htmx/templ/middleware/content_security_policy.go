package middleware

import (
	"math/rand"
	"mmyoungman/templ/utils"
	"net/http"
)

func ContentSecurityPolicy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		nonce := GenerateNonce()

		utils.SetContextValue(r, utils.CspNonceCtxKey, nonce)

		w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' 'nonce-"+
			nonce+
			"'")

		// @MarkFix Add reporting endpoint?
		//w.Header().Set("Content-Security-Policy", "default-src 'self'; report-to csp-endpoint")
		//w.Header().Set("Reporting-Endpoints", "csp-endpoint=\"/csp-reports\"")

		next.ServeHTTP(w, r)
	})
}

func GenerateNonce() string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	newNonce := make([]rune, 20)
	for i := 0; i < 20; i++ {
		newNonce[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(newNonce)
}
