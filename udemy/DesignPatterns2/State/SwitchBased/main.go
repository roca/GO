package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Switch is a simple state machine

// Combination lock

type State int

const (
	Locked State = iota
	Failed
	Unlocked
)

func main() {
	code := "1234"
	state := Locked
	entry := strings.Builder{}

	for {
		switch state {
		case Locked:
			r, _, _ := bufio.NewReader(os.Stdin).ReadRune()
			entry.WriteRune(r)

			if entry.String() == code {
				state = Unlocked
				break
			}

			if strings.Index(code, entry.String()) != 0 {
				state = Failed
			}
		case Failed:
			fmt.Println("FAILED")
			entry.Reset()
			state = Locked
		case Unlocked:
			fmt.Println("UNLOCKED")
			return
		}
	}

}
