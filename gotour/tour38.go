package main

import (
	"code.google.com/p/go-tour/pic"
	//"math"
)

func Pic(dx, dy int) (pic [][]uint8) {

	pic = make([][]uint8, dx)
	for i := range pic {
		pic[i] = make([]uint8, dy)
	}

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			pic[x][y] = uint8(x * y)
		}
	}
	return

}

func main() {
	pic.Show(Pic)
}
