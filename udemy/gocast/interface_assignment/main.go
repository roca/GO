package main

import (
	"fmt"
)

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func (s square) printArea() {
	fmt.Println(s.getArea)
}

type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
func (t triangle) printArea() {
	fmt.Println(t.getArea)
}

type shape interface {
	printArea()
}
