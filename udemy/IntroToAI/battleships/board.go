package main

var shipTypes = []struct {
	name string
	size int
}{
	{"Carrier", 5},
	{"Battleship", 4},
	{"Cruiser", 3},
	{"Submarine", 3},
	{"Destroyer", 2},
}

type Board [boardSize][boardSize]string

type Position struct {
	row, col int
}
