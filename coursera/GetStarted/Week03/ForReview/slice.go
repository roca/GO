package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var s string
	sInt := make([]int, 3)

	for {
		fmt.Print("Enter number: ")
		fmt.Scan(&s)

		if s == "X" {
			break
		}

		i, _ := strconv.Atoi(s)
		sInt = append(sInt, i)

		sort.Ints(sInt)
		fmt.Println(sInt)
	}
}
