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
	mux.Post("/auth", app.authenticate)
	mux.Post("/refresh", app.refresh)

	// test handler

	// JWT protected routes
	mux.Route("/users", func(mux chi.Router) {
		// use auth middleware

		mux.Get("/", app.allUser)
		mux.Get("/{userID}", app.getUser)
		mux.Delete("/{userID}", app.deleteUser)
		mux.Put("/", app.insertUser)
		mux.Patch("/", app.updateUser)
	})

	return mux
}
