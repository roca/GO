package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

var animals = map[string]Animal{"cow": {"grass", "walk", "moo"},
	"bird":  {"worms", "fly", "peep"},
	"snake": {"mice", "slither", "hsss"}}

var empty = Animal{}

func main() {
	log.SetOutput(os.Stdout)
	fmt.Println("For exit print \"exit\"")
	promt()
}

func promt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")
		if scanner.Scan() {
			query := scanner.Text()
			if query == "exit" {
				return
			}

			animal, request := parse(query)

			switch request {
			case "eat":
				Eat(animals[animal])
			case "move":
				Move(animals[animal])
			case "speak":
				Speak(animals[animal])
			default:
				fmt.Println("Unknown command")

			}

		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
			return
		}
	}

}

func parse(line string) (string, string) {

	name := strings.SplitN(line, " ", 2)
	name = append(name, "")

	return strings.ToLower(name[0]), strings.ToLower(name[1])
}

func Eat(animal Animal) {
	if animal != empty {
		fmt.Println(animal.food)
	} else {
		fmt.Println("Unknown animal")
	}

}

func Move(animal Animal) {
	if animal != empty {
		fmt.Println(animal.locomotion)
	} else {
		fmt.Println("Unknown animal")
	}
}

func Speak(animal Animal) {
	if animal != empty {
		fmt.Println(animal.noise)
	} else {
		fmt.Println("Unknown animal")
	}
}
