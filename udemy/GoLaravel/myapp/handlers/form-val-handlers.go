package handlers

import (
	"fmt"
	"myapp/data"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

func (h *Handlers) Form(w http.ResponseWriter, r *http.Request) {
	vars := make(jet.VarMap)
	validator := h.App.Validator(nil)
	vars.Set("validator", validator)
	vars.Set("user", data.User{})

	err := h.App.Render.Page(w, r, "form", vars, nil)
	if err != nil {
		h.App.ErrorLog.Println(err)
	}
}

func (h *Handlers) PostForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.App.ErrorLog.Println(err)
		return
	}

	validator := h.App.Validator(nil)

	validator.Required(r, "first_name", "last_name", "email")
	validator.IsEmail("email", r.PostForm.Get("email"))

	validator.Check(len(r.Form.Get("first_name")) > 1, "first_name", "Must be at least 2 characters")
	validator.Check(len(r.Form.Get("last_name")) > 1, "last_name", "Must be at least 2 characters")

	if !validator.Valid() {
		vars := make(jet.VarMap)
		vars.Set("validator", validator)
		vars.Set("user", data.User{
			FirstName: r.Form.Get("first_name"),
			LastName:  r.Form.Get("last_name"),
			Email:     r.Form.Get("email"),
		})

		if err := h.App.Render.Page(w, r, "form", vars, nil); err != nil {
			h.App.ErrorLog.Println(err)
			return
		}
		return
	}
	fmt.Fprint(w, "valid data")
}
