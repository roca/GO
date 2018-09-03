package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

const resp = `<html>
    <head>
        <title>Simple Web App</title>
	</head>
    <body>
        <h1>Simple Web App</h1>
        <p>Hello World : Added GOCV to {{.Hostname}} xxy</p>
    </body>
</html>`

type Data struct {
	Hostname string // has to be uppercase/exportable/public
}

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	substitute := Data{hostname}
	tmpl, err := template.New("resp").Parse(resp)
	err = tmpl.Execute(w, substitute)

}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
