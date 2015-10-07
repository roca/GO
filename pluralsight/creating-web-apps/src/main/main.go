package main

import (
	"net/http"
	"os"
	"text/template"
)

func main() {

	templates := populateTemplates()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		requestedFile := req.URL.Path[1:]
		template := templates.Lookup(requestedFile + ".html")
		if template != nil {
			template.Execute(w, nil)
		} else {
			w.WriteHeader(404)
		}
	})

	http.ListenAndServe(":8000", nil)
}

func populateTemplates() *template.Template {
	result := template.New("templates")

	basePath := "templates"
	tempateFolder, _ := os.Open(basePath)
	defer tempateFolder.Close()

	templatePathsRaw, _ := tempateFolder.Readdir(-1)

	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		}
	}

	result.ParseFiles(*templatePaths...)

	return result
}

// func main() {
// 	http.Handle("/", new(MyHandler))
// 	http.ListenAndServe(":8000", nil)
// }

// type MyHandler struct {
// 	http.Handler
// }
//
// func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//
// 	path := "public/" + req.URL.Path
//
// 	if data, err := ioutil.ReadFile(string(path)); err == nil {
// 		var contentType string
// 		if strings.HasSuffix(path, ".css") {
// 			contentType = "text/css"
// 		} else if strings.HasSuffix(path, ".html") {
// 			contentType = "text/html"
// 		} else if strings.HasSuffix(path, ".js") {
// 			contentType = "application/javascript"
// 		} else if strings.HasSuffix(path, ".png") {
// 			contentType = "image/png"
// 		} else {
// 			contentType = "text/plain"
// 		}
// 		w.Header().Add("Content Type", contentType)
// 		w.Write(data)
// 	} else {
// 		w.WriteHeader(404)
// 		w.Write([]byte("404 - " + http.StatusText(404)))
// 	}
// }
//
// const doc = `
// {{template "header" .Title}}
// 		<body>
// 		<h1> List of Fruit </h1>
// 		<ul>
// 		    {{range .Fruit}}
// 				<li> {{.}} </li>
// 				{{end}}
// 		</ul>
// 		</body>
// {{template "footer"}}
// `
//
// const header = `
// <!DOCTYPE html>
// <html>
//     <head><title>{{.}}</title></head>
// `
//
// const footer = `
// </html>
// `
//
// type Context struct {
// 	Fruit [3]string
// 	Title string
// }
