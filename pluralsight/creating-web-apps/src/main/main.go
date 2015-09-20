package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":8000", nil)
}

type MyHandler struct {
	http.Handler
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "public/html" + req.URL.Path
	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}
