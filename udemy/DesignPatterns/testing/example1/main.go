package main

import "os"


func main() {
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])
	result := sum(a,b)
	fmt.Printf("The sum of %d and %d is %d\n",a,b,result)
}

fun sum(a,b int) int {
	return a + b
}