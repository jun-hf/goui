package handlers

import (
	"log/slog"
	"net/http"
)

type HandlerHTTP func(w http.ResponseWriter, r *http.Request) error 

func Make(h HandlerHTTP) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("Error", "handler", err, "path", r.URL.Path)
		}
	}
}