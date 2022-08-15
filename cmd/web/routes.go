package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jacstn/golang-url-shortner/pkg/handlers"
)

func routes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(LoadSession)
	mux.Get("/", handlers.Home)
	mux.Get("/about", handlers.About)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
