// File name: ...\s07\10_pointers_slice\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

// ASSIGNMENT - in previous sections we learned that a slice is a pointer to its
// underlying array. Create an array of strings and a slice of string.
// Using pointers prove that the slice uses the same address region of the array.
func main() {

	days := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	fmt.Printf("days=%v\n", days)
	fmt.Printf("len(days)=%d\n", len(days))
	fmt.Printf("cap(days)=%d\n\n", cap(days))

	weekdays := days[3:6]
	fmt.Printf("weekdays=%v\n", weekdays)
	fmt.Printf("len(weekdays)=%d\n", len(weekdays))
	fmt.Printf("cap(weekdays)=%d\n\n", cap(weekdays))

	fmt.Printf("&days[3]=%x\n", &days[3])
	fmt.Printf("&weekdays[0]=%x\n", &weekdays[0])

	fmt.Println()
	fmt.Printf("&days[5]=%x\n", &days[5])
	fmt.Printf("&weekdays[2]=%x\n", &weekdays[2])
}
