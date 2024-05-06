package main

import (
	"fmt"
	"go-breaders/pets"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/roca/go-toolkit/v2"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.New("dog"))
}

func (app *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolkit.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.New("cat"))
}
