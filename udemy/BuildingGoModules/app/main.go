package main

import (
	"fmt"
	"net/http"

	"github.com/roca/GO/tree/staging/udemy/BuildingGoModules/toolkit"
)

func main() {
	var tools toolkit.Tools

	http.HandleFunc("/upload", HandleFileUpload)

	fmt.Println(tools.RandomString(10))

	http.ListenAndServe(":8080", nil)
}

func HandleFileUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File upload request received")
}
