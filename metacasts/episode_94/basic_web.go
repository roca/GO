// basic_web
package main

import (
	//"fmt"
	"html/template"
	//"io/ioutil"
	"github.com/elazarl/go-bindata-assetfs"
	"log"
	"net/http"
)

type HomePageHandler struct{}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	renderTemplate("templates/index.html", w)
}

func renderTemplate(tmplName string, w http.ResponseWriter) {
	//data, _ := ioutil.ReadFile(tmplName)
	data, _ := Asset(tmplName)
	tmpl, err := template.New("Index").Parse(string(data))
	if err != nil {

		log.Panic(err)
	}

	tmpl.Execute(w, nil)
}

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", &HomePageHandler{})
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(&assetfs.AssetFS{Asset, AssetDir, "assets"})))

	return mux
}

func main() {

	log.Println(AssetNames)

	http.ListenAndServe(":3000", NewMux())
}
