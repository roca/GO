// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/9_MSdcdlNQ

// Declare a struct type named Point with two fields, X and Y of type int.
// Implement a factory function for this type and a method that accept a value
// of this type and calculates the distance between the two points. What is
// the nature of this type?
package main

import (
	"fmt"
	"math"
)

// Add imports.

// Declare struct type named Point that represents a point in space.
type Point struct {
	X int
	Y int
}

// Declare a function named New that returns a Point based on X and Y
// positions on a graph.
func New(x int, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

// Declare a method named Distance that finds the length of the hypotenuse
// between two points. Pass one point in and return the answer.
// Forumula is the square root of (x2 - x1)^2 + (y2 - y1)^2
// Use the math.Pow and math.Sqrt functions.
func (p Point) Distance(p2 Point) float64 {
	first := math.Pow(float64(p2.X-p.X), 2)
	second := math.Pow(float64(p2.Y-p.Y), 2)
	return math.Sqrt(first + second)
}

// main is the entry point for the application.
func main() {
	// Declare the first point.

	// Declare the second point.

	// Calculate the distance and display the result.
	p1 := New(37, -76)
	p2 := New(26, -80)

	dist := p1.Distance(p2)
	fmt.Println("Distance", dist)
}
