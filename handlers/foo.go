package handlers

import (
	"net/http"

	"github.com/jun-hf/goui/views/foo"
)

func HandlerFoo(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, foo.Foo())
}