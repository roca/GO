package main

import (
	"log"
	"net/http"

	"github.com/roca/go-toolkit/v2"
)

func main() {
	mux := routes()

	log.Println("Starting application on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))

	mux.HandleFunc("/api/login", login)
	mux.HandleFunc("/api/logout", logout)

	return mux
}

func login(w http.ResponseWriter, r *http.Request) {
	var tools toolkit.Tools

	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := tools.ReadJSON(w, r, &payload)
	if err != nil {
		tools.ErrorJSON(w, err)
		return
	}

	var resPayLoad toolkit.JSONResponse

	if payload.Email == "me@here.com" && payload.Password == "password" {
		resPayLoad.Error = false
		resPayLoad.Message = "You have been logged in"
		_ = tools.WriteJSON(w, http.StatusAccepted, resPayLoad)
		return
	}

	resPayLoad.Error = true
	resPayLoad.Message = "Invalid credentials"
	_ = tools.WriteJSON(w, http.StatusUnauthorized, resPayLoad)

}

func logout(w http.ResponseWriter, r *http.Request) {
	var tools toolkit.Tools

	payload := toolkit.JSONResponse{
		Message: "You have been logged out",
	}

	_ = tools.WriteJSON(w, http.StatusAccepted, payload)
}
