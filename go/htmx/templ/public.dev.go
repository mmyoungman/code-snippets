//go:build !prod

package main

import (
	"fmt"
	"net/http"
	"os"
)

func public() http.Handler { // Cannot go:embed public like in public.prod.go if we want @hotreload
	fmt.Println("building static files for development")
	return http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public")))
}
