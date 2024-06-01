package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/jun-hf/goui/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("PLEASE provide .env file")
	}
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		log.Fatal("PLEASE provide listenAddr in .env")
	}
	router := chi.NewMux()
	router.Handle("/*", public())
	router.Get("/foo", handlers.Make(handlers.HandlerHome))

	slog.Info("HTTP server started", "listenAddr", listenAddr)
	http.ListenAndServe(listenAddr, router)

	fmt.Println("Hello world!")
}	