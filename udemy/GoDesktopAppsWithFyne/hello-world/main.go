package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	output *widget.Label
}

var myApp App

func main() {
	a := app.New()
	w := a.NewWindow("Hello, World")

	output, entry, btn := myApp.makeUI()

	w.SetContent(container.NewVBox(output, entry, btn))
	w.Resize(fyne.NewSize(500, 500))
	w.ShowAndRun()
}

func (app *App) makeUI() (*widget.Label, *widget.Entry, *widget.Button) {
	output := widget.NewLabel("Hello, World!")
	entry := widget.NewEntry()
	btn := widget.NewButton("Enter", func() {
		app.output.SetText(entry.Text)
	})
	btn.Importance = widget.HighImportance

	app.output = output

	return output, entry, btn
}
