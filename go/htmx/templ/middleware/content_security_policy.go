package middleware

import (
	"mmyoungman/templ/utils"
	"net/http"
)

func ContentSecurityPolicy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		nonce := utils.GenerateRandomStr()

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
