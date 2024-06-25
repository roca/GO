package main

import "fmt"

// Item is an item in the game
type Item struct {
	X int
	Y int
}

const (
	maxX = 1000
	maxY = 600
)


func NewItem(x, y int) (*Item, error) {
	if x < 0 || x >= maxX || y < 0 || y >= maxY {
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}

	item := Item{
		X: x,
		Y: y,
	}

	// The Go compiler does "escape analysis" to determine if the item variable can be allocated on the stack or heap
	// returning a pointer to the item variable will force the compiler to allocate the item variable on the heap
	return &item, nil
}


// if you want to mutate, use a pointer receiver
func (i *Item) Move(x,y int) {
	i.X = x
	i.Y = y
}



func main() {
	var i1 Item
	fmt.Println(i1)
	fmt.Printf("i1: %#v\n", i1)

	i2 := Item{1, 2}
	fmt.Printf("i2: %#v\n", i2)

	i3 := Item{
		Y: 20,
		//X: 10,
	}
	fmt.Printf("i3: %#v\n", i3)

	fmt.Println(NewItem(10,20))
	fmt.Println(NewItem(10,-20))

	i3.Move(100,200)
	fmt.Printf("i3 (Move): %#v\n", i3)

}

