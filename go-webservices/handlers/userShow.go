package handlers

import (
	"html/template"
	"net/http"
)

//UserNew form to into a new User
func UserShow(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage("show")
	if err != nil {
		p = &Page{Title: "show"}
	}
	t, _ := template.ParseFiles("./views/users/show.html")
	t.Execute(w, p)
}
