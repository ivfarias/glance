package handler

import (
	"net/http"

	"github.com/glanceapp/glance/internal/glance"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.Handle("/", glance.Start(glance.Options{
		ConfigDir: "config",
		AssetsDir: "assets",
	}))
	mux.ServeHTTP(w, r)
}
