package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/roca/go-toolkit/v2"
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
	mux.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		var payload = struct {
			Message string `json:"message"`
		}{
			Message: "Hello World",
		}
		var tools toolkit.Tools
		_ = tools.WriteJSON(w, http.StatusOK, payload, nil)
	})

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
