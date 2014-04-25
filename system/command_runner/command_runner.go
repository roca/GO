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
	//fmt.Println("Working on command: ", *command)
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("Done with ", *command)
	done <- true
}

func main() {

	commands := make([]Command, 10000)

	for i := range commands {
		commands[i] = Command(fmt.Sprintf("command %d", i))
	}

	done := make(chan bool)
	defer func() { close(done) }()

	for i := range commands {
		go commands[i].execute(done)
	}

	all_done := len(commands)
	doneCount := 0
	timeout := time.After(6 * time.Second)

	for doneCount != all_done {

		select {
		case <-done:
			doneCount++
		case <-timeout:
			fmt.Println("timed out!")
			doneCount = all_done
		}

	}

	fmt.Println("all done !")
}
