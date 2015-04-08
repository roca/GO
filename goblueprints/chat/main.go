package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(`
			<html>
				<head>
					<title>Chat</title>
				</head>
				<body>
					Let's char!
				</body>
			</html>
		`))

}

func main() {
	http.HandleFunc("/", handler)
	//start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
