// martini
package main

import (
	"github.com/codegangsta/martini"
	"github.com/russross/blackfriday"
	//"log"
	"net/http"
	//"net/http/cgi"
)

func main() {
	m := martini.Classic()

	m.Post("/go-bin/martini.cgi", func(r *http.Request) []byte {

		body := r.FormValue("body")
		return blackfriday.MarkdownBasic([]byte(body))
	})

	m.Get("/go-bin/martini.cgi", func(r *http.Request) []byte {

		body := r.FormValue("body")
		return blackfriday.MarkdownBasic([]byte(body))
	})
	m.Run()

	//http.HandleFunc("/go-bin/quadratic.cgi", m.ServeHTTP)
	//if err := cgi.Serve(http.HandlerFunc(m.ServeHTTP)); err != nil {
	//	log.Fatal("failed to start server", err)
	//}

}
