package v1

import "github.com/GOCODE/go-webservices/api"

func API() {
	api.Version = 1
	api.StartServer()

