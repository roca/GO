package main

import (
	"errors"
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

func (m *Matrix[T]) At(row, col int) T {
	pos := (row * m.Cols) + col
	if pos > len(m.data) {
		return 0
	}

	return m.data[pos]
}

func (m *Matrix[T]) Set(row, col int, value T) error {
	pos := (row * m.Cols) + col
	if pos > len(m.data) {
		return fmt.Errorf("Out of bounds")
	}

	m.data[pos] = value

	return nil
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

	m, err := NewMatrix[float64](10, 3)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("m:", m)

	m.Set(3, 2, 5)
	fmt.Println("m[3,2]:", m.At(3, 2))

	fmt.Println(Max([]int{3, 1, 2}))     // 3, nil
	fmt.Println(Max([]float64{3, 1, 2})) // 3, nil
	fmt.Println(Max[int](nil))           // 0, Max of empty silce

}

func Max[T Number](values []T) (T, error) {
	if values == nil {
		var zero T
		return zero, errors.New("Max of empty silce")
	}

	max := values[0]
	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}

	return max, nil
}
