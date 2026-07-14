package shapes

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"

	"example.com/strategy"
)

type ImageSquare struct {
	strategy.PrintOutPut
	DestinationFilePath string
}

func (t *ImageSquare) Print() error {
	width := 800
	height := 600

	origin := image.Point{0, 0}

	bgImage := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: width, Y: height},
	})

	bgColor := image.Uniform{color.RGBA{R: 70, G: 70, B: 70, A: 0}}
	quality := &jpeg.Options{Quality: 75}

	draw.Draw(bgImage, bgImage.Bounds(), &bgColor, origin, draw.Src)

	squareWidth := 200
	squareHeight := 200

	squareColor := image.Uniform{color.RGBA{R: 255, G: 0, B: 0, A: 1}}
	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (width / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})
	squareImage := image.NewRGBA(square)

	draw.Draw(bgImage, squareImage.Bounds(), &squareColor, origin, draw.Src)

	if t.Writer == nil {
		return fmt.Errorf("No writer stored on ImageSquare")
	}

	if err := jpeg.Encode(t.Writer, bgImage, quality); err != nil {
		return fmt.Errorf("Error writing image to disk: %g", err)
	}

	if t.LogWriter != nil {
		r := bytes.NewReader([]byte("Image written in provided writer\n"))
		io.Copy(t.LogWriter, r)
	}
	return nil
}
