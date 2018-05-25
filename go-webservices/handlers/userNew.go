package handlers

import (
	"html/template"
	"net/http"
)

//UserNew form to into a new User
func UserNew(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage("new")
	if err != nil {
		p = &Page{Title: "new"}
	}
	t, _ := template.ParseFiles("./views/users/new.html")
	t.Execute(w, p)
}
