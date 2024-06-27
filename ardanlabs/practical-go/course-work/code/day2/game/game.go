package main

import (
	"fmt"
	"slices"
)

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
	// go build -gcflags=-m shows escape analysis
	return &item, nil
}

// if you want to mutate, use a pointer receiver
func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

type Player struct {
	Name string
	// X int
	Item // Embed Item in Player
	Keys []Key
}

// func (p *Player) Move(x, y int) {
// 	p.Item.Move(x, y)

// }

//Rule of thumb: Accept interfaces, return concrete types

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}

}

type mover interface {
	Move(x, y int)
}

type Key byte

// Go's version of enums
const (
	Jade Key = iota + 1
	Copper
	Crystal
	invalidKey // internal (not exported)
)

// Implement the fmt.Stringer interface
func (k Key) String() string {
	switch k {
	case Jade:
		return "Jade"
	case Copper:
		return "Copper"
	case Crystal:
		return "Crystal"
	default:
		return fmt.Sprintf("Unknown key: %d", k)
	}
}

// FoundKey adds a key to the player's keyring
// If the key is not a known key, it returns an error
// If the key is already in the keyring, it does nothing
func (p *Player) FoundKey(k Key) error {
	if k < Jade || k >= invalidKey {
		return fmt.Errorf("%s", k)
	}

	// if !containsKey(p.Keys, k) {
	// 	p.Keys = append(p.Keys, k)
	// }

	if !slices.Contains(p.Keys, k) {
		p.Keys = append(p.Keys, k)
	}

	return nil
}

func containsKey(keys []Key, k Key) bool {
	for _, key := range keys {
		if key == k {
			return true
		}
	}
	return false
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

	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(10, -20))

	i3.Move(100, 200)
	fmt.Printf("i3 (Move): %#v\n", i3)

	p1 := Player{
		Name: "Parzival",
		Item: Item{X: 500, Y: 300},
	}
	fmt.Printf("p1: %#v\n", p1)
	fmt.Printf("p1.X: %#v\n", p1.X)
	fmt.Printf("p1.Item.X: %#v\n", p1.Item.X)
	p1.Move(400, 600)
	fmt.Printf("p1 (Move): %#v\n", p1)

	ms := []mover{
		&i1,
		&p1,
		&i2,
	}
	moveAll(ms, 0, 0)
	for _, m := range ms {
		fmt.Printf("%#v\n", m)
	}

	k := Copper
	fmt.Printf("k: %d\n", k)
	fmt.Printf("k: %s\n", k)
	fmt.Printf("k: %#v\n", k)

	fmt.Println("key:", Key(17))

	p1.FoundKey(Jade)
	fmt.Printf("p1: %v\n", p1)
}
