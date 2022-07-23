package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/koopa0/ws/internal/handlers"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/chat", handlers.Home)
	mux.Get("/ws", handlers.WsEndpoint)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
