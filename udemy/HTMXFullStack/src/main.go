package main

import (
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseGlob("./templates/*.html"))

		data := struct {
			Name        string
			Title       string
			Description string
			Socials     map[string]string
		}{
			Name:        "Alice",
			Title:       "Visitor",
			Description: "This paragraph is the new body of our page",
			Socials: map[string]string{
				"Facebook":  "example.com",
				"Twitter/X": "@example.com",
				"Instagram": "example-pics",
				"LikedIn":   "example-inc",
			},
		}

		err := tmpl.ExecuteTemplate(w, "home.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":3000", nil)
}
