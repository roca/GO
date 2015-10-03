package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content Type", "text/html")
		templates := template.New("template")
		templates.New("test").Parse(doc)
		templates.New("header").Parse(header)
		templates.New("footer").Parse(footer)

		context := Context{
			Fruit: [3]string{"Lemon", "Orange", "Apple"},
			Title: "Fruits",
		}
		templates.Lookup("test").Execute(w, context)

	})

	http.ListenAndServe(":8000", nil)
}

// func main() {
// 	http.Handle("/", new(MyHandler))
// 	http.ListenAndServe(":8000", nil)
// }

type MyHandler struct {
	http.Handler
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	path := "public/" + req.URL.Path

	if data, err := ioutil.ReadFile(string(path)); err == nil {
		var contentType string
		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else {
			contentType = "text/plain"
		}
		w.Header().Add("Content Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}

const doc = `
{{template "header" .Title}}
		<body>
		<h1> List of Fruit </h1>
		<ul>
		    {{range .Fruit}}
				<li> {{.}} </li>
				{{end}}
		</ul>
		</body>
{{template "footer"}}
`

const header = `
<!DOCTYPE html>
<html>
    <head><title>{{.}}</title></head>
`

const footer = `
</html>
`

type Context struct {
	Fruit [3]string
	Title string
}
