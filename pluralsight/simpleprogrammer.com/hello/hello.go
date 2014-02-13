package main

import (
	"./greeting"
	"fmt"
)

func RenameToFrog(r greeting.Renamable) {
	r.Rename("Frog")
}

func main() {

	salutaions := greeting.Salutations{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "What is up?"},
	}

	fmt.Fprintf(&salutaions[0], "The count is %d", 10)

	done := make(chan bool, 2)

	go func() {
		salutaions.Greet(greeting.CreatePrintFunction("<C>"), true, 6)
		done <- true
		done <- true
		fmt.Println("Done")
	}()
	salutaions.Greet(greeting.CreatePrintFunction("?"), true, 6)
	<-done

}
