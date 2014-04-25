// command_runner
package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

type Command string

type Runner interface {
	run(done chan bool)
}

func (command *Command) run(done chan bool) {
	//fmt.Println("--------------------------------Working on command: ", *command)

	out, err := exec.Command(string(*command)).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output from %s is\n%s\n", *command, out)

	//time.Sleep(5000 * time.Millisecond)
	//fmt.Println("--------------------------------Done with ", *command)
	done <- true
}

func main() {

	commands := []Command{"ls", "date", "pwd"}

	//commands := make([]Command, 10000)

	//for i := range commands {
	//	commands[i] = Command(fmt.Sprintf("command %d", i))
	//}

	done := make(chan bool)
	defer func() { close(done) }()

	for i := range commands {
		go commands[i].run(done)
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
