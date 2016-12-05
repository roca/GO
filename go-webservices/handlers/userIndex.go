package handlers

import (
	"html/template"
	"net/http"
)

//UserIndex show list of users and a form to add more
func UserIndex(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage("index")
	if err != nil {
		p = &Page{Title: "index"}
	}
	t, _ := template.ParseFiles("./views/users/index.html")
	t.Execute(w, p)
}
