package main

import (
	"fmt"
)

func truncFloatToInt(num float64) int {
	return int(num)
}

func main() {
	var input1 float64
	fmt.Printf("Enter First Float Number:")
	fmt.Scanf("%f", &input1)
	firstInput := truncFloatToInt(input1)
	println("Integer:", firstInput)
	var input2 float64
	fmt.Printf("\n\nEnter Second Float Number:")
	fmt.Scanf("%f", &input2)
	secondInput := truncFloatToInt(input2)
	println("Integer:", secondInput)
}
