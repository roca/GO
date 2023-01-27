package main

import "fmt"

// Example of a Virtual Proxy Pattern

type Image interface {
	Draw()
}


type Bitmap struct {
	filename string
}

func NewBitmap(filename string) *Bitmap {
	fmt.Println("Loading image from", filename)
	return &Bitmap{filename: filename}
}

func (b *Bitmap) Draw() {
	// Draw the bitmap
	fmt.Println("Drawing image", b.filename)
}

func DrawImage(image Image) {
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Done drawing the image")
}

type LazyBitmap struct {
	filename string
	bitmap   *Bitmap
}

func (lb *LazyBitmap) Draw() {
	if lb.bitmap == nil {
		lb.bitmap = NewBitmap(lb.filename)
	}
	lb.bitmap.Draw()
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{filename: filename}
}

func main() {
	bmp := NewLazyBitmap("image.png")
	DrawImage(bmp)
	fmt.Println("")
	DrawImage(bmp)
}