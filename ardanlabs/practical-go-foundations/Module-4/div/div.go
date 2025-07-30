package main

import "fmt"

func main() {
	fmt.Println(safeDiv(7, 3))
	fmt.Println(safeDiv(7, 0))
}

func safeDiv(a, b int) (q int, err error) {
	defer func() {
		if e := recover(); e != nil {
			//fmt.Println("ERROR:", e)
			err = fmt.Errorf("%v", e)
		}
	}()

	q = div(a, b)

	return
}

func div(a, b int) int {
	/*
		if b == 0 {
			panic("division by zero")
		}
	*/
	return a / b
}
