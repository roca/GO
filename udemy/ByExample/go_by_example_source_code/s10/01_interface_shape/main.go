// File name: ...\s10\01_interface_shape\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
type rectangle struct {
	w, l int //width & length
}

func (c *rectangle) area() int {
	return c.w * c.l
}

func (c *rectangle) perim() int {
	return 2 * (c.l + c.w)
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
type square struct {
	s int //side
}

func (c *square) area() int {
	return c.s * c.s
}

func (c *square) perim() int {
	return 4 * c.s
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
type shape interface {
	area() int
	perim() int
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
func info(s shape) {
	fmt.Printf("area()=%d perim()=%d\n", s.area(), s.perim())
}

func totalArea(shapes ...shape) int {
	var totalArea int
	for _, s := range shapes {
		totalArea += s.area()
	}
	return totalArea
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
func main() {

	r1 := rectangle{2, 3}
	fmt.Printf("r1.area()=%d\n", r1.area())
	fmt.Printf("r1.perim()=%d\n", r1.perim())

	r2 := rectangle{1, 4}
	fmt.Printf("r2.area()=%d\n", r2.area())
	fmt.Printf("r2.perim()=%d\n", r2.perim())

	s1 := square{3}
	fmt.Printf("s1.area()=%d\n", s1.area())
	fmt.Printf("s1.perim()=%d\n", s1.perim())

	s2 := square{7}
	fmt.Printf("s2.area()=%d\n", s2.area())
	fmt.Printf("s2.perim()=%d\n", s2.perim())

	fmt.Println("\n= = = = = = = = =")
	info(&r1)
	info(&r2)
	info(&s1)
	info(&s2)

	fmt.Println("\n= = = = = = = = =")
	fmt.Printf("Total Area=%d\n", totalArea(&r1, &r2, &s1, &s2))

	var shapes [4]shape
	shapes[0] = &r1
	shapes[1] = &r2
	shapes[2] = &s1
	shapes[3] = &s2

	totalArea := 0
	for _, s := range shapes {
		totalArea += s.area()
	}

	fmt.Printf("Total Area=%d\n", totalArea)

	info(shapes[0])
}
