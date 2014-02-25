package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w int
	h int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel

}
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)

}
func (img Image) At(x, y int) color.Color {

	v := uint8(x * y)

	return color.RGBA{v, v, 255, 255}

}

func main() {
	m := Image{1000, 1000}
	pic.ShowImage(m)
}
