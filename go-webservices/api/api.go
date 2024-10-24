package api

import (
	"net/http"
	"os"

	"github.com/GOCODE/go-webservices/handlers"

	"github.com/gorilla/mux"
)

var Version int

func StartServer() {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/{user:[0-9]+}", handlers.Hello)
	//Views
	gorillaRoute.HandleFunc("/users", handlers.UserIndex).Methods("GET")
	gorillaRoute.HandleFunc("/users/new", handlers.UserNew).Methods("GET")
	gorillaRoute.HandleFunc("/users/{id:[0-9]+}", handlers.UserShow).Methods("GET")
	gorillaRoute.HandleFunc("/users/{id:[0-9]+}/edit", handlers.UserEdit).Methods("GET")
	//API
	gorillaRoute.HandleFunc("/api/users", handlers.UserCreate).Methods("POST")
	gorillaRoute.HandleFunc("/api/users", handlers.UsersRetrieve).Methods("GET")

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))

	http.Handle("/", gorillaRoute)
	port := os.Getenv("PORT")

	http.ListenAndServe(":"+port, nil)
}
