package main

import (
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseGlob("./templates/*.html"))

		err := tmpl.ExecuteTemplate(w, "greetingFragment", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":3000", nil)
}
