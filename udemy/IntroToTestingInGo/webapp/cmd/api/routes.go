package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// register middleware
	mux.Use(middleware.Recoverer)
	// mus.Use(app.enableCORS)

	// authentication routes - auth handler, refresh

	// test handler

	// JWT protected routes

	return mux
}
