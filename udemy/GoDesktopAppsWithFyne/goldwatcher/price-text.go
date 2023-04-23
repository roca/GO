package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2"
)

func (app *Config) getGoldText() (*canvas.Text, *canvas.Text, *canvas.Text) {
	var g Gold
	var open, current, change *canvas.Text

	gold, err := g.GetPrices()
	if err != nil {
		grey := color.NRGBA{R: 155, G: 155, B: 155, A: 255}
		open = canvas.NewText("Open: Unreachable", grey)
		current = canvas.NewText("Current: Unreachable", grey)
		change = canvas.NewText("Change: Unreachable", grey)
		return open, current, change
	}

	displayColor := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	if gold.Price < gold.PreviousClose {
		displayColor = color.NRGBA{R: 180, G: 0, B: 0, A: 255}
	}

	openTxt := fmt.Sprintf("Open: $%.4f %s", gold.PreviousClose, gold.Currency)
	currentTxt := fmt.Sprintf("Current: $%.4f %s", gold.Price, gold.Currency)
	changeTxt := fmt.Sprintf("Change: $%.4f %s", gold.Change, gold.Currency)

	open = canvas.NewText(openTxt, nil)
	current = canvas.NewText(currentTxt, displayColor)
	change = canvas.NewText(changeTxt, displayColor)

	open.Alignment = fyne.TextAlignLeading
	current.Alignment = fyne.TextAlignCenter
	change.Alignment = fyne.TextAlignTrailing

	return open, current, change
}
