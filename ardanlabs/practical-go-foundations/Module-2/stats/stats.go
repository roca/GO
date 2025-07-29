package main

import (
	"fmt"
	"time"
)

func NewMatrix[T Number](rows, cols int) (*Matrix[T], error) {
	if rows <= 0 || cols <= 0 {
		return nil, fmt.Errorf("bad dimensions: %d/%d", rows, cols)
	}

	m := Matrix[T]{
		Rows: rows,
		Cols: cols,
		data: make([]T, rows*cols),
	}

	return &m, nil
}

type Matrix[T Number] struct {
	Rows int
	Cols int

	data []T
}

type Number interface {
	~int | ~float64
}

func Relu[T Number](i T) T {
	if i < 0 {
		return 0
	}

	return i
}

func main() {
	fmt.Println(Relu(7))
	fmt.Println(Relu(-1))
	fmt.Println(Relu(1.2))
	fmt.Println(Relu(time.February))

	m,err := NewMatrix[float64](10,3)
	if err != nil {
		fmt.Println("ERROR:",err)
		return
	}

	fmt.Println("m:",m)

}
