package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// Constant

const cellSize = 60

// Variables for color
var (
	green     = color.RGBA{G: 255, A: 255}
	darkGreen = color.RGBA{R: 0, G: 100, B: 32, A: 255}
	red       = color.RGBA{R: 255, A: 255}
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
	}

	upLeft := image.Point{}
	lowRight := image.Point{X: width, Y: height}

	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	draw.Draw(img, img.Bounds(), &image.Uniform{C: color.Black}, image.Point{}, draw.Src)

	// draw squares on the image
	for i, row := range g.Walls {
		for j, col := range row {
			p := Point{
				Row: i,
				Col: j,
			}

			if col.wall {
				// draw black square for wall
				g.drawSquare(col, p, img, color.Black, cellSize, j*cellSize, i*cellSize)
			} else if g.inSolution(p) {
				// draw green square for solution path
				g.drawSquare(col, p, img, green, cellSize, j*cellSize, i*cellSize)
			} else if col.State.Row == g.Start.Row && col.State.Col == g.Start.Col {
				// draw darkGreen square for start
				g.drawSquare(col, p, img, darkGreen, cellSize, j*cellSize, i*cellSize)
			} else if col.State.Row == g.Goal.Row && col.State.Col == g.Goal.Col {
				// draw red square for goal
				g.drawSquare(col, p, img, red, cellSize, j*cellSize, i*cellSize)
			} else if col.State == g.CurrentNode.State {
				// draw orange square for current node
				g.drawSquare(col, p, img, orange, cellSize, j*cellSize, i*cellSize)
			} else if inExplored(Point{i, j}, g.Explored) {
				// draw yellow square for explored
				g.drawSquare(col, p, img, yellow, cellSize, j*cellSize, i*cellSize)
			} else {
				// draw white square for unvisited
				g.drawSquare(col, p, img, color.White, cellSize, j*cellSize, i*cellSize)
			}
		}
	}

	f, _ := os.Create(outFile)
	_ = png.Encode(f, img)
}

// drawSquare
func (g *Maze) drawSquare(w Wall, p Point, img *image.RGBA, fillColor color.Color, size, offsetX, offsetY int) {
	patch := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.Draw(patch, patch.Bounds(), &image.Uniform{C: fillColor}, image.Point{}, draw.Src)

	if !w.wall {
		// Print the x y coordinates of this cell
		g.printLocation(p, color.Black, patch)
	}

	draw.Draw(img, image.Rect(offsetX, offsetY, offsetX+size, offsetY+size), patch, image.Point{}, draw.Src)
}

// printLocation
func (g *Maze) printLocation(p Point, c color.Color, patch *image.RGBA) {
	point := fixed.Point26_6{X: fixed.I(6), Y: fixed.I(40)}
	d := &font.Drawer{
		Dst:  patch,
		Src:  image.NewUniform(c),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(fmt.Sprintf("[%d,%d]", p.Row, p.Col))
}
