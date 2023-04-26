package main

import (
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type Config struct {
	App                 fyne.App
	InfoLog             *log.Logger
	ErrorLog            *log.Logger
	MainWindow          fyne.Window
	PriceContainer      *fyne.Container
	ToolBar             *widget.Toolbar
	PriceChartContainer *fyne.Container
	HttpClient          *http.Client
}

var myApp Config

func main() {
	// create a fyne application
	fyneApp := app.NewWithID("desertfoxdev.org.goldwatcher.preferences")
	myApp.App = fyneApp
	myApp.HttpClient = &http.Client{}

	//create our loggers
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// open a connection to our local database. SQLite

	// create a database repository using the Repo interface

	// create and size a fyne window
	myApp.MainWindow = fyneApp.NewWindow("GoldWatcher")
	myApp.MainWindow.Resize(fyne.NewSize(770, 410))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.CenterOnScreen()
	myApp.MainWindow.SetMaster()

	myApp.makeUI()

	// show and run the application
	myApp.MainWindow.ShowAndRun()

}
