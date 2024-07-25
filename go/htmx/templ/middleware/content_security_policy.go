package middleware

import (
	"context"
	"mmyoungman/templ/utils"
	"net/http"
)

func ContentSecurityPolicy(nonce string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			newCtx := context.WithValue(ctx, utils.CspNonceCtxKey, nonce)
			*r = *r.WithContext(newCtx)

			w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' 'nonce-" + nonce + "'; script-src 'self' 'nonce-" + nonce + "'")

			// @MarkFix Add reporting endpoint?
			//w.Header().Set("Content-Security-Policy", "default-src 'self'; report-to csp-endpoint")
			//w.Header().Set("Reporting-Endpoints", "csp-endpoint=\"/csp-reports\"")

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
