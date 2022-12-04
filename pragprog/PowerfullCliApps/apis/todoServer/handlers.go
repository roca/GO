package main

import "net/http"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	content := "There's an API here"
	replyTextContent(w, r, http.StatusOK, content)
}
