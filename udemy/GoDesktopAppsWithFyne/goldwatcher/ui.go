package main

import "fyne.io/fyne/v2/container"

func (app *Config) makeUI() {

	// get the current price of gold
	// opening-price, current-price, price-change
	openingPrice, currentPrice, priceChange := app.getGoldText()
	// https://data-asg.goldprice.org/dbxRates/USD

	// put price information into a container
	priceContent := container.NewGridWithColumns(3,
		openingPrice,
		currentPrice,
		priceChange,
	)

	app.PriceContainer = priceContent

	// add container to the window
	finalContent := container.NewVBox(priceContent)

	app.MainWindow.SetContent(finalContent)
}
