package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// Image URL: https://goldprice.org/charts/gold_3d_b_o_usd_x.png

func (app *Config) pricesTab() *fyne.Container {
	chart := app.getChart()
	charContainer := container.NewVBox(chart)
	app.PriceChartContainer = charContainer

	return charContainer
}

func (app *Config) getChart() *canvas.Image {
	apiURL := fmt.Sprintf("https://goldprice.org/charts/gold_3d_b_o_%s_x.png", strings.ToLower(currency))
	var img *canvas.Image

	err := app.downloadFile(apiURL, "gold.png")
	if err != nil {
		// use bundled image
		img = canvas.NewImageFromResource(resourceUnreachablePng)

	} else {
		img = canvas.NewImageFromFile("gold.png")
	}

	img.SetMinSize(fyne.Size{
		Width:  400,
		Height: 400,
	})

	img.FillMode = canvas.ImageFillContain

	return img
}

func (app *Config) downloadFile(URL, filename string) error {
	// get the response bytes from calling a url
	response, err := app.HttpClient.Get(URL)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New("received wrong response code when downloading image")
	}
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return err
	}

	out, err := os.Create(fmt.Sprintf("./%s", filename))
	if err != nil {
		return err
	}
	err = png.Encode(out, img)
	if err != nil {
		return err
	}
	defer out.Close()

	return nil
}
