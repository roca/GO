package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	greet := widget.NewLabel("Hello World!")
	input := widget.NewEntry()
	input.SetPlaceHolder("World")

	w.SetContent(container.NewVBox(
		greet,
		input,
		widget.NewButton("Greet", func() {
			message := "Hello " + input.Text + "!"
			greet.SetText(message)
		}),
	))

	w.ShowAndRun()
}
