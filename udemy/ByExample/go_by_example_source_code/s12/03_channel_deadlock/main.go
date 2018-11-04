// File name: ...\s12\03_channel_deadlock\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

func main() {
	c := make(chan string)
	c <- "No one likes my channel!" //fatal error: all goroutines are asleep - deadlock!
}
