package main

import (
	"fmt"
)

func main() {
	var i Item
	fmt.Printf("i: %#v\n", i)

	i = Item{10, 20} // Musr specify all fields
	fmt.Printf("i: %#v\n", i)

	// Can be in any order, can omit fields
	i = Item{
		Y: 22,
		// X: 11,
	}
	fmt.Printf("i: %#v\n", i)


	fmt.Println(New(10,20))
	fmt.Println(New(10,2000))

	/* Aside: %#v for debugging/logging
	a, b := 1, "1"
	fmt.Printf("a=%v, b=%v\n",a,b)
	fmt.Printf("a=%#v, b=%#v\n",a,b)
	*/
}

/* Possible factory funcs
func New(x,y int) Item
func New(x,y int) *Item
func New(x,y int) (Item, error)
func New(x,y int) (*Item, error)
*/

func New(x, y int) (Item, error) {
	if x < 0 ||
		x > maxX ||
		y < 0 ||
		y > maxY {
		return Item{}, fmt.Errorf("%d/%d out of max range %d/%d", x, y, maxX, maxY)
	}

	return Item{x, y}, nil
}

const (
	maxX = 600
	maxY = 400
)

type Item struct {
	X int
	Y int
}
