package main

import (
	"html/template"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseGlob("./templates/*.html"))

		data := struct {
			Name        string
			Title       string
			Description string
			Socials     map[string]string
			Features    []string
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
			Features: []string{
				"Customizable Products",
				"24/7 Customer Support",
				"Reliable and Secure",
			},
		}

		err := tmpl.ExecuteTemplate(w, "home.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/functions", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseGlob("./templates/*.html"))

		// numberString := r.PathValue("number")
		// number, err := strconv.Atoi(numberString)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }

		data := struct {
			Name        string
			CurrentDate time.Time
			Number      int
			Items       []string
		}{
			Name:        "John Doe",
			CurrentDate: time.Now(),
			Number:      7,
			Items:       []string{"Apples", "oranges", "Bananas"},
		}
		data.Number = 15

		err := tmpl.ExecuteTemplate(w, "functions.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":3000", nil)
}
