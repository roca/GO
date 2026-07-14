package main

import (
	"fmt"
	"sort"
)

func main() {
	cart := []string{"apple", "orange", "banana"}
	fmt.Println("len:", len(cart))
	fmt.Println("cart[1]:", cart[1])

	for i, c := range cart {
		fmt.Println(i, c, cart[i])
	}

	cart = append(cart, "milk")
	fmt.Println("cart:", cart)

	fruit := cart[:3]
	fmt.Println("fruit:", fruit)

	fruit = append(fruit, "lemon")
	fmt.Println("fruit:", fruit)
	fmt.Println("cart:", cart) // "milk" replaced by "lemon"

	out := concat([]string{"A", "B"}, []string{"C"})
	fmt.Println("concat:", out) // [A, B, C]

	values := []float64{3, 1, 2}
	fmt.Println(median(values)) // 2

	values = []float64{3, 1, 2, 4}
	fmt.Println(median(values)) // 2.5
	fmt.Println(values)

	values = []float64{1, 2}
	fmt.Println(median(values)) // 1.5

	players := []Player{
		{"Rick", 10_000},
		{"Morty", 11},
	}

	// Value semantic "for" loop
	for _, p := range players {
		p.Score += 100
	}

	// Pointer semantic "for" loop
	for i := range players {
		players[i].Score += 100
	}

	fmt.Println(players)
}

type Player struct {
	Name  string
	Score int
}

func median(ns []float64) float64 {

	// Copy in order not to mutate input params
	vals := make([]float64, len(ns))
	copy(vals, ns)

	sort.Float64s(vals)
	i := len(vals) / 2

	if len(vals)%2 != 0 {
		return vals[i]
	}
	return (vals[i-1] + vals[i]) / 2
}

func concat(s1, s2 []string) []string {
	size_cap := len(s1) + len(s2)
	s := make([]string, size_cap)

	copy(s, s1)
	copy(s[len(s1):], s2)

	return s
}
