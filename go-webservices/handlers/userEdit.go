package handlers

import (
	"html/template"
	"net/http"
)

//UserNew form to into a new User
func UserEdit(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage("edit")
	if err != nil {
		p = &Page{Title: "edit"}
	}
	t, _ := template.ParseFiles("./views/users/edit.html")
	t.Execute(w, p)
}
