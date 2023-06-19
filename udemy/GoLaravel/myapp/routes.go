package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	// middle must come before any routes

	// add routes here
	a.App.Routes.Get("/", a.Handlers.Home)
	a.App.Routes.Get("/go-page", a.Handlers.GoPage)
	a.App.Routes.Get("/jet-page", a.Handlers.JetPage)
	a.App.Routes.Get("/sessions", a.Handlers.SessionsTest)

	a.App.Routes.Get("/users/login", a.Handlers.UserLogin)
	a.App.Routes.Post("/users/login", a.Handlers.PostUserLogin)
	a.App.Routes.Get("/users/logout", a.Handlers.Logout)

	a.App.Routes.Get("/create-user", a.Handlers.CreateUser)
	a.App.Routes.Get("/get-user/{id}", a.Handlers.GetUserByID)
	a.App.Routes.Get("/update-user/{id}", a.Handlers.UpdateUserByID)

	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
