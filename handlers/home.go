package handlers

import (
	"net/http"

	"github.com/jun-hf/goui/views/home"
)

func HandlerHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Home())
}