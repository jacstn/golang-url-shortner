package main

import (
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

	return mux
}
