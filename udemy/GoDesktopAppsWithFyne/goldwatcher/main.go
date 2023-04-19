package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Config struct {
	App      fyne.App
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

var myApp Config

func main() {
	// create a fyne application
	fyneApp := app.NewWithID("desertfoxdev.org.goldwatcher.preferences")

	//create our loggers

	// open a connection to our local database. SQLite

	// create a database repository using the Repo interface

	// create and size a fyne window

	// show and run the application

}
