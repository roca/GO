// basic_web
package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type HomePageHandler struct{}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	renderTemplate("templates/index.html", w)
}

func renderTemplate(tmplName string, w http.ResponseWriter) {
	data, _ := ioutil.ReadFile(tmplName)
	tmpl, err := template.New("Index").Parse(string(data))
	if err != nil {

		log.Panic(err)
	}

	tmpl.Execute(w, nil)
}

func main() {
	http.Handle("/", &HomePageHandler{})

	barHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Bar")
	}

	http.HandleFunc("/bar", barHandler)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":3000", nil)
}
