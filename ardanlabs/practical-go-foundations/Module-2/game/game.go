package main

import (
	"fmt"
	"slices"
)

func main() {
	var i Item
	fmt.Printf("i: %#v\n", i)

	i = Item{10, 20} // Musr specify all fields
	fmt.Printf("i: %#v\n", i)

	// Can be in any order, can omit fields
	// i = Item{
	// 	Y: 22,
	// 	// X: 11,
	// }
	// fmt.Printf("i: %#v\n", i)

	fmt.Println(New(10, 20))
	fmt.Println(New(10, 2000))

	/* Aside: %#v for debugging/logging
	a, b := 1, "1"
	fmt.Printf("a=%v, b=%v\n",a,b)
	fmt.Printf("a=%#v, b=%#v\n",a,b)
	*/

	i.Move(10, 20)
	fmt.Printf("i (after move): %#v\n", i)

	p1 := Player{
		Name: "Parzival",
	}
	fmt.Printf("p1: %+v\n", p1)
	fmt.Printf("p1.X: %+v\n", p1.Item.X)

	p1.Move(100, 200)
	fmt.Printf("p1: %+v\n", p1)

	fmt.Println(p1.Found(Copper))
	fmt.Println(p1.Found(Copper))
	fmt.Println(p1.Found(Key(7)))
	fmt.Println("keys:", p1.Keys)

	ms := []Mover{
		&i,
		&p1,
	}

	moveAll(ms, 50, 50)
	for _, m := range ms {
		fmt.Println(m)
	}

}

/*

	func Sort(s (s Sortable) {

	// ...
	}

	type Sortable interface {
	  Less(i, j int) bool
	  Swap(i, j int)
	  Len() int
	}

*/

type Mover interface {
	Move(int, int)
}

func moveAll(ms []Mover, dx, dy int) {
	for _, m := range ms {
		m.Move(dx, dy)
	}
}

// Move moves i by delta x & delta y
func (i *Item) Move(dx, dy int) {
	i.X += dx
	i.Y += dy
}

/* Possible factory funcs
func New(x,y int) Item
func New(x,y int) *Item
func New(x,y int) (Item, error)
func New(x,y int) (*Item, error)
*/

func New(x, y int) (*Item, error) {
	if x < 0 ||
		x > maxX ||
		y < 0 ||
		y > maxY {
		return nil, fmt.Errorf("%d/%d out of max range %d/%d", x, y, maxX, maxY)
	}

	i := Item{x, y}

	// Do: go build -gcflags=-m
	// You'll see './game.go:47:2: moved to heap: i'
	return &i, nil
}

const (
	maxX = 600
	maxY = 400
)

type Item struct {
	X int
	Y int
}

type Player struct {
	Name string
	Keys []Key
	Item
}

// func (p *Player) AddKey(key string) {
// 	if ok, err := p.Found(key); !ok && err != nil {
// 		p.Keys = append(p.Keys, key)
// 	}
// }

// func (p Player) Found(key string) (bool, error) {

// 	// if slices.ContainsFunc(p.Keys, func(k string) bool {
// 	// 	return k == key
// 	// }) {
// 	// 	return true, nil
// 	// }

// 	if slices.Contains(p.Keys, key) {
// 		return true, nil
// 	}
// 	return false, fmt.Errorf("unknown key: %#v", key)
// }

type Key byte

func (k Key) String() string {
	switch k {
	case Copper:
		return "copper"
	case Jade:
		return "jade"
	case Crystal:
		return "crystal"
	default:
		return fmt.Sprintf("<Key %d>", k)
	}
}

const (
	Copper Key = iota + 1
	Jade
	Crystal
)

func (p *Player) Found(key Key) error {
	switch key {
	case Copper, Jade, Crystal:
		// OK
	default:
		return fmt.Errorf("unknown key: %v", key)
	}

	if !slices.Contains(p.Keys, key) {
		p.Keys = append(p.Keys, key)
	}

	return nil
}
