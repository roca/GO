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

	c := make(chan greeting.Salutation)
	c2 := make(chan greeting.Salutation)
	go salutaions.ChannelGreeter(c)
	go salutaions.ChannelGreeter(c2)

	for {
		select {
		case s, ok := <-c:
			if ok {

				fmt.Println(s, ":1")
			} else {
				return
			}

		case s, ok := <-c2:
			if ok {
				fmt.Println(s, ":2")
			} else {
				return
			}
		}
	}

}
