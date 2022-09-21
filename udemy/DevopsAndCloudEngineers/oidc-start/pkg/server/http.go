package server

import (
	"fmt"
	"net/http"
)

type server struct {
	PrivateKey []byte
	Config     Config
	LoginRequests map[string]LoginRequest
}

func newServer(privateKey []byte, config Config) *server {
	return &server{
		PrivateKey: privateKey,
		Config:     config,
		LoginRequests: make(map[string]LoginRequest),
	}
}

func Start(httpServer *http.Server, privateKey []byte, config Config) error {
	s := newServer(privateKey, config)

	http.HandleFunc("/authorization", s.authorization)
	http.HandleFunc("/token", s.token)
	http.HandleFunc("/login", s.login)
	http.HandleFunc("/jwks.json", s.jwks)
	http.HandleFunc("/.well-known/openid-configuration", s.discovery)
	http.HandleFunc("/userinfo", s.userinfo)

	return httpServer.ListenAndServe()
}

func returnError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
	fmt.Printf("Error: %s\n", err)
}
