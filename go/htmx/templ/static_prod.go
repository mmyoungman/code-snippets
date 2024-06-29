//go:build prod

package main

import (
	"embed"
	"fmt"
	"net/http"
	"os"
)

//go:embed public
var publicFS embed.FS

func public() http.Handler {
	return http.FileServerFS(publicFS)
}