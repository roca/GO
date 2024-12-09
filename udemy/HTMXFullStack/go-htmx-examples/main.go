package main

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Uppercase a string
func Upper(str string) string {
	return strings.ToUpper(str)
}

// Format a date
func FmtDate(t time.Time) string {
	return t.Format("January 2, 2006")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
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
}

func FunctionsHandler(w http.ResponseWriter, r *http.Request) {

	funcMap := template.FuncMap{
		"Upper":    Upper,
		"FmtDate": FmtDate,
	}
	funcTmpl := template.New("funcMap").Funcs(funcMap)

	tmpl := template.Must(funcTmpl.ParseFiles("./functions.html"))

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
	numberString := r.PathValue("number")
	if numberString != "" {
		number, _ := strconv.Atoi(numberString)
		data.Number = number
	}

	err := tmpl.ExecuteTemplate(w, "functions.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	http.HandleFunc("/", HomeHandler)

	http.HandleFunc("GET /functions/{number}", FunctionsHandler)
	http.HandleFunc("GET /functions", FunctionsHandler)

	http.ListenAndServe(":3000", nil)
}
