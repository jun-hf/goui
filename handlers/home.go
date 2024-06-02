package handlers

import (
	"net/http"

	"github.com/jun-hf/goui/views/home"
)

func HandlerHome(w http.ResponseWriter, r *http.Request) error {
	var a = make(map[string]bool)
	a["hell"] = true
	return Render(w, r, home.Home(a))
}