// Author: Kaustubh Olpadkar
// New Animals Go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{ name string }

func (cow Cow) Eat()   { fmt.Println("grass") }
func (cow Cow) Move()  { fmt.Println("walk") }
func (cow Cow) Speak() { fmt.Println("moo") }

type Bird struct{ name string }

func (bird Bird) Eat()   { fmt.Println("worms") }
func (bird Bird) Move()  { fmt.Println("fly") }
func (bird Bird) Speak() { fmt.Println("peep") }

type Snake struct{ name string }

func (snake Snake) Eat()   { fmt.Println("mice") }
func (snake Snake) Move()  { fmt.Println("slither") }
func (snake Snake) Speak() { fmt.Println("hsss") }

func Prompt() {
	fmt.Print("> ")
}

var reader = bufio.NewReader(os.Stdin)

func GetInput() (string, string, string) {
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Invalid User Input")
		os.Exit(1)
	}

	input = strings.TrimSuffix(input, "\n")
	inputs := strings.Split(input, " ")

	if len(inputs) != 3 {
		fmt.Println("Invalid User Input - Exactly Three Arguments Expected")
		os.Exit(1)
	}

	return inputs[0], inputs[1], inputs[2]
}

func ExecuteQuery(name, action string, animal Animal) {
	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Invalid Action", action)
		os.Exit(1)
	}
}

var animalMap = map[string]func(string) Animal{
	"cow":   func(name string) Animal { return Cow{name} },
	"bird":  func(name string) Animal { return Bird{name} },
	"snake": func(name string) Animal { return Snake{name} },
}

func main() {
	animals := map[string]Animal{}
	for {
		Prompt()
		command, name, action := GetInput()

		switch command {
		case "newanimal":
			if action == "cow" || action == "bird" || action == "snake" {
				animals[name] = animalMap[action](name)
			} else {
				fmt.Println("Invalid Animal Type", action)
				os.Exit(1)
			}
		case "query":
			animal := animals[name]
			ExecuteQuery(name, action, animal)
		}
	}
}
