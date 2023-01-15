package main

import (
	"log"
	"net/http"

	"github.com/roca/GO/tree/staging/udemy/BuildingGoModules/toolkit"
)

func main() {
	mux := routes()

	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func routes() http.Handler {

	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/download", downloadFile)

	return mux
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
	t := toolkit.Tools{}
	t.DownloadStaticFile(w, r, "./files", "pic.jpg", "puppy.jpg")
}
