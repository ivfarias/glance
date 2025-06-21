package main

import (
	"net/http"

	"github.com/glanceapp/glance/internal/glance"
)

// Handler is the entrypoint for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.Handle("/", glance.Start(glance.Options{
		ConfigDir: "config",
		AssetsDir: "assets",
	}))
	mux.ServeHTTP(w, r)
}
