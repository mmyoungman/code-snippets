//go:build prod

package main

import (
	"embed"
	"net/http"
)

//go:embed public
var publicFS embed.FS

func public() http.Handler { // @MarkFix check compile executable actually works when moved somewhere else without associated files - including migrations
	return http.FileServerFS(publicFS)
}
