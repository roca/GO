package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func main() {
	m := mat64.NewDense(3, 5, nil)

	// fill in m with some elements
	for i := 0; i < 3; i++ {
		m.Set(i, i, 99)
	}

	// wrong way to print matrix
	fmt.Println("m : ", m)

	// proper way
	// refer to https://godoc.org/github.com/gonum/matrix/mat64#Excerpt
	fmt.Printf("m :\n%v\n\n", mat64.Formatted(m, mat64.Prefix(" "), mat64.Excerpt(2)))

	// print all m elements
	fmt.Printf("m :\n%v\n\n", mat64.Formatted(m, mat64.Prefix(""), mat64.Excerpt(0)))

	data := []float64{1, 2, 3}
	m2 := mat64.NewDense(3, 1, data)

	// print all m2 elements
	fmt.Printf("m2 :\n%v\n\n", mat64.Formatted(m2, mat64.Prefix(""), mat64.Excerpt(0)))

	// to build this matrix
	// ⎡1  2  3⎤
	// ⎢4  5  6⎥
	// ⎣7  8  9⎦

	// use SetRow or SetCol
	m3 := mat64.NewDense(3, 3, nil)
	m3.SetRow(0, data)

	data2 := []float64{4, 5, 6}
	m3.SetRow(1, data2)

	data3 := []float64{7, 8, 9}
	m3.SetRow(2, data3)

	// print all m3 elements
	fmt.Printf("m3 :\n%v\n\n", mat64.Formatted(m3, mat64.Prefix(""), mat64.Excerpt(0)))

	// get transpose with m3.T()
	fmt.Printf("m3 transpose :\n%v\n\n", mat64.Formatted(m3.T(), mat64.Prefix(""), mat64.Excerpt(0)))

}
