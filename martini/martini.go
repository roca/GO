// martini
package main

import (
	"github.com/codegangsta/martini"
	"github.com/russross/blackfriday"
	"net/http"
	//"os"
)

func main() {
	m := martini.Classic()

	m.Post("/go-bin/martini.cgi", func(r *http.Request) []byte {

		body := r.FormValue("body")
		return blackfriday.MarkdownBasic([]byte(body))
	})

	//os.Setenv("PORT", "80")

	//m.Run()
}
