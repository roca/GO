package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getFloatParam(scanner *bufio.Scanner, parameter string) float64 {
	fmt.Printf("Enter the %s as a number:\n", parameter)
	scanner.Scan()
	param, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		panic("Please enter a valid number")
	}

	return param
}

// GenDisplaceFn takes three float64 arguments and returns a function
// to calculate displacement at time t
func GenDisplaceFn(acc, velo, displ float64) func(float64) float64 {
	return func(t float64) float64 {
		// displacement: 0.5a*t^2 + v*t + s
		return (0.5 * acc * t * t) + (velo * t) + displ
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	acc := getFloatParam(scanner, "acceleration")
	velo := getFloatParam(scanner, "initial velocity")
	displ := getFloatParam(scanner, "initial displacement")
	fn := GenDisplaceFn(acc, velo, displ)
	t := getFloatParam(scanner, "time")

	fmt.Println("The displacement after", t, "seconds is:", fn(t))
}
