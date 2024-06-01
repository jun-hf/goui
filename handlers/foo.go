package handlers

import (
	"errors"
	"net/http"
)

func HandlerFoo(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("FOO"))
	return errors.New("foo error")
}