// command_runner
package main

import (
	"fmt"
	"time"
)

type Command string

type Executable interface {
	execute(done chan bool)
}

func (command *Command) execute(done chan bool) {
	fmt.Println("Working on command: ", *command)
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Done with command!", *command)
	done <- true
}

func main() {

	var a_command Command
	a_command = "test"
	fmt.Println("Hello World!")
	done := make(chan bool)
	go a_command.execute(done)

	for {
		select {
		case <-done:
			break
		}
	}
}
