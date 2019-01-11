package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal : type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat method should print the animal’s food
func (a Animal) Eat() {
	fmt.Println(a.food)
}

// Move method should print the animal’s locomotion
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

// Speak method should print the animal’s spoken sound
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

// Request take in a request and executes the approriate action
func (a Animal) Request(request string) {
	switch request {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Println("This animal doesn't know how to do that! It can only eat, move or speak")
	}
}

func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	animals := map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}

	// Scanf wont work if your input has spaces :)
	consoleReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		userInput, _ := consoleReader.ReadString('\n')
		userInput = strings.TrimSuffix(userInput, "\n")

		animalRequest := strings.Split(userInput, " ")
		animal, request := animalRequest[0], animalRequest[1]

		animals[animal].Request(request)
	}

}
