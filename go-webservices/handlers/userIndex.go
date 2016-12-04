package handlers

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	filename := title + ".html"
	body, err := ioutil.ReadFile("./views/users/" + filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

//UserIndex show list of users and a form to add more
func UserIndex(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage("index")
	if err != nil {
		p = &Page{Title: "index"}
	}
	t, _ := template.ParseFiles("./views/users/index.html")
	t.Execute(w, p)
}
