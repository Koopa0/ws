package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/koopa0/ws/internal/handlers"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/home", handlers.Home)

	return mux
}
