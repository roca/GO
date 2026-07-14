package main

import (
	"fmt"
	"log"
)

func main() {
	// fmt.Println(div(1, 0))
	v,err := safeDiv(1, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(v)
	fmt.Println(safeDiv(7, 2))
}

// named return values
func safeDiv(a, b int) (q int, err error) {
	// q & err are local variables in safeDiv
	// (just like a & b)
	defer func() {
		// e's type is interface{} *not* error
		if e := recover(); e != nil {
			log.Println("ERROR:", e)
			err = fmt.Errorf("%#v", e)
		}
	}()
	// panic("I'm panicking!")

	/* Miki don't like naked returns 
	q = a / b
	return
	*/
	return a / b, nil
}

func div(a, b int) int {
	return a / b
}
