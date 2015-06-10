// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/5KL4HipSJ-

// Declare a struct type named Point with two fields, X and Y of type int.
// Implement a factory function for this type and a method that accept a value
// of this type and calculates the distance between the two points. What is
// the nature of this type.
package main

import (
	"fmt"
	"math"
)

// Point represents a point in space.
type Point struct {
	X int
	Y int
}

// New returns a Point based on X and Y positions on a graph.
func New(x int, y int) Point {
	return Point{x, y}
}

// Distance finds the length of the hypotenuse between two points.
// Forumula is the square root of (x2 - x1)^2 + (y2 - y1)^2
func (p Point) Distance(p2 Point) float64 {
	first := math.Pow(float64(p2.X-p.X), 2)
	second := math.Pow(float64(p2.Y-p.Y), 2)
	return math.Sqrt(first + second)
}

// main is the entry point for the application.
func main() {
	p1 := New(37, -76)
	p2 := New(26, -80)

	dist := p1.Distance(p2)
	fmt.Println("Distance", dist)
}
