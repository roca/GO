package main

import (
	"fmt"

	"github.com/GOCODE/pluralsight/go-testing/code4/03_CheckEqual/src/pack"
)

func main() {
	fmt.Println(pack.QuickSort(7, 3, 9, 1))
	fmt.Println(pack.QuickSort(9, 8, 7, 6))
}
