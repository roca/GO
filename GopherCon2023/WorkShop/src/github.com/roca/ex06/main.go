package main

import "fmt"

// func printString(s string) { fmt.Println(s) }

// func printInt(i int) { fmt.Println(i) }

// func printBool(b bool) { fmt.Println(b) }

type printableTypes interface {
	string | int | bool
}

// func myPrint[T printableTypes](t T) { fmt.Println(t) }
func myPrintAny[T any](t T) { fmt.Println(t) }

func main() {
	myPrintAny("hello")
	myPrintAny(5)
	myPrintAny(false)
}
