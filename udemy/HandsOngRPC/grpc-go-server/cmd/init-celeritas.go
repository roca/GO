package main

import (
	"log"
	"grpc-go-server/data"

	"os"

	"github.com/roca/celeritas"
)

type application struct {
	App        *celeritas.Celeritas
	Models     data.Models
}

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init celertias
	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "grpc-go-server"

	app := &application{
		App:        cel,
	}

	app.Models = data.New(app.App.DB.Pool)

	return app

}
