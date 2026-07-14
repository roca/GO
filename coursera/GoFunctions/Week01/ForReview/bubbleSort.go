package main

import "fmt"

func main() {
	var n int
	var err error
	slice := make([]int, 0)

	fmt.Print("Enter count of numbers (from 1 to 10): ")
	_, err = fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Error while scanning count of numbers: ", err)
		return
	}
	if n > 10 || n < 1 {
		fmt.Println("Invalid count value: ", n)
		return
	}

	for i := 1; i <= n; i++ {
		var number int

		fmt.Print("Enter number ", i, ": ")
		_, err = fmt.Scanln(&number)
		if err != nil {
			fmt.Println("Error while scanning number: ", err)
			return
		}

		slice = append(slice, number)
	}

	BubbleSort(slice)

	fmt.Print("Sorted data: ")
	for _, value := range slice {
		fmt.Print(value, " ")
	}

	fmt.Println()
}

func BubbleSort(values []int) {
	for i := len(values) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if values[j] > values[j+1] {
				Swap(values, j)
			}
		}
	}
}

func Swap(slice []int, index int) {
	slice[index], slice[index+1] = slice[index+1], slice[index]
}
