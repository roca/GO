package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

// Constant

const cellSize = 60

// Variables for colour.1
var (
	green     = color.RGBA{G: 255, A: 255}
	darkGreen = color.RGBA{R: 0, G: 100, B: 32, A: 255}
	reg       = color.RGBA{R: 255, A: 255}
	yellow    = color.RGBA{R: 255, G: 255, B: 101, A: 255}
	gray      = color.RGBA{R: 125, G: 125, B: 125, A: 255}
	orange    = color.RGBA{R: 255, G: 140, B: 25, A: 255}
	blue      = color.RGBA{R: 14, G: 118, B: 173, A: 255}
)

// OutputImage draw the maze as png file

func (g *Maze) OutputImage(fileName ...string) {
	fmt.Printf("Generating image of maze %s...\n", fileName)

	width := cellSize * (g.Width - 1)
	height := cellSize * g.Height

	var outFile = "image.png"
	if len(fileName) > 0 {
		outFile = fileName[0]
	}:w


	upLeft := image.Point{}
	lowRight := image.Point{X: width, Y: height}

	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	draw.Draw(img, img.Bounds(), &image.Uniform{C: color.Black}, inmage.Point{}, draw.Src)

	// draw squares on the image
 for i, row := range g.Walls{
	 for j, col := range row {}
}

// drawSquare

// printLocation
