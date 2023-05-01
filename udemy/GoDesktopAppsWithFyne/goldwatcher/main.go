package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"goldwatcher/repository"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	_ "github.com/glebarez/go-sqlite"
)

type Config struct {
	App                 fyne.App
	InfoLog             *log.Logger
	ErrorLog            *log.Logger
	DB                  repository.Repository
	MainWindow          fyne.Window
	PriceContainer      *fyne.Container
	ToolBar             *widget.Toolbar
	PriceChartContainer *fyne.Container
	HttpClient          *http.Client
}

func main() {
	var myApp Config

	// create a fyne application
	fyneApp := app.NewWithID("desertfoxdev.org.goldwatcher.preferences")
	myApp.App = fyneApp
	myApp.HttpClient = &http.Client{}

	//create our loggers
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// open a connection to our local database. SQLite
	sqlDb, err := myApp.connectSQL()
	if err != nil {
		myApp.ErrorLog.Fatal(err)
	}

	// create a database repository using the Repo interface
	myApp.setupDB(sqlDb)

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

func (app *Config) connectSQL() (*sql.DB, error) {
	path := ""
	if os.Getenv("DB_PATH") != "" {
		path = os.Getenv("DB_PATH")
	} else {
		path = app.App.Storage().RootURI().Path() + "/sql.db"
	}
	app.InfoLog.Println("DB Path: ", path)

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *Config) setupDB(sqlDB *sql.DB) {
	app.DB = repository.NewSQLiteRepository(sqlDB)

	err := app.DB.Migrate()
	if err != nil {
		app.ErrorLog.Println(err)
		app.ErrorLog.Fatal(err)
	}
}
