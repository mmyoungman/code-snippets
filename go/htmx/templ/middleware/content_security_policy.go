package middleware

import (
	"net/http"
)

func ContentSecurityPolicy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'")

		// @MarkFix Add reporting endpoint?
		//w.Header().Set("Content-Security-Policy", "default-src 'self'; report-to csp-endpoint")
		//w.Header().Set("Reporting-Endpoints", "csp-endpoint=\"/csp-reports\"")

		next.ServeHTTP(w, r)
	})
}
