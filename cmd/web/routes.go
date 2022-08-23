package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jacstn/golang-url-shortner/internal/handlers"
)

func routes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(WriteToConsole)
	mux.Use(LoadSession)
	mux.Use(NoSurf)
	mux.Get("/", handlers.Home)
	mux.Get("/about", handlers.About)
	mux.Get("/list-all", handlers.List)
	mux.Post("/new-url", handlers.CreateUrl)
	mux.Get("/view-url/{id}", handlers.ViewUrl)
	mux.Get("/{id}", handlers.Redirect)
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
