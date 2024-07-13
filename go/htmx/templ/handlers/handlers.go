package handlers

import (
	"log/slog"
	"net/http"
)

type HTTPHandler func(writer http.ResponseWriter, request *http.Request) error

func Make(handler HTTPHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			slog.Error("HTTP handler error",
				"err", err,
				"path", request.URL.Path)
		}
	}
}
