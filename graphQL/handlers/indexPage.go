package handlers

import (
	"html/template"
	"net/http"
)

//IndexPage show list of users and a form to add more
func IndexPage(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage("index")
	if err != nil {
		p = &Page{Title: "GraphiQL"}
	}
	t, _ := template.ParseFiles("./views/index.html")
	t.Execute(w, p)
}
