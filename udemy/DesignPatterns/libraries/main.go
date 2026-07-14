package main

import (
	"fmt"
	"github.com/roca/GO/udemy/DesignPatterns/libraries/arithemitic"
)

func main() {
	sumRes := arithmetic.Sum(5, 6)
	subRes := arithmetic.Subtract(10, 5)
	multiplyRes := arithmetic.Multiply(8, 7)
	divideRes,_ := arithmetic.Divide(10, 2)
	fmt.Printf("5+6 is %d, 10-5 is %d, 8*7 is %d and 10/2 is %f\n", sumRes, subRes, multiplyRes, divideRes)
}
