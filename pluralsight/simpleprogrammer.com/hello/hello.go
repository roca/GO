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

	fmt.Fprintf(&salutaions[0],"The count is %d",10)

	salutaions.Greet(greeting.CreatePrintFunction("?"), true, 6)
}
