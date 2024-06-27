package main

import (
	"fmt"
	"log"
)

func main() {
	// fmt.Println(div(1, 0))
	fmt.Println(safeDiv(1, 0))
	fmt.Println(safeDiv(7, 2))
}

// named return values
func safeDiv(a, b int) (q int, err error) {
	// q & err are local variables in safeDiv
	// (just like a & b)
	defer func() {
		if e := recover(); e != nil {
			log.Println("ERROR:", e)
			err = fmt.Errorf("%v", e)
		}
	}()

	/* Miki don't like this kind of programming
	q = a / b
	return
	*/
	return a / b, nil
}

func div(a, b int) int {
	return a / b
}
