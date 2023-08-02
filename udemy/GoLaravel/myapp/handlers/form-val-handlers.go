package handlers

import (
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
