package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// Animal ...
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow ...
type Cow struct{}

// Eat : "grass"
func (a Cow) Eat() {
	fmt.Println("grass")
}

// Move : walk
func (a Cow) Move() {
	fmt.Println("walk")
}

// Speak : "moo"
func (a Cow) Speak() {
	fmt.Println("moo")
}

// Bird ...
type Bird struct{}

// Eat : "worms"
func (a Bird) Eat() {
	fmt.Println("worms")
}

// Move : "fly"
func (a Bird) Move() {
	fmt.Println("fly")
}

// Speak : "peep"
func (a Bird) Speak() {
	fmt.Println("peep")
}

// Snake ...
type Snake struct{}

// Eat : "mice"
func (a Snake) Eat() {
	fmt.Println("mice")
}

// Move : "slither"
func (a Snake) Move() {
	fmt.Println("slither")
}

// Speak : "hsss"
func (a Snake) Speak() {
	fmt.Println("hsss")
}

var animals map[string]Animal

// Request take in a request and executes the approriate action
func Request(a Animal, request string) {
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

func animalFactory(animalType string) (Animal, error) {
	switch animalType {
	case "cow":
		return Cow{}, nil
	case "bird":
		return Bird{}, nil
	case "snake":
		return Snake{}, nil
	}
	return nil, errors.New("Can't create animal of this type")
}

func main() {

	animals = map[string]Animal{}

	// Scanf wont work if your input has spaces :)
	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("Three inputs are required. (Example: 'newanimal timmy bird' or 'query timmy speak'")

	for {
		fmt.Print(">")
		userInput, _ := consoleReader.ReadString('\n')
		userInput = strings.TrimSuffix(userInput, "\n")

		animalRequest := strings.Split(userInput, " ")
		if len(animalRequest) < 3 {
			fmt.Println("Three inputs are required. (Example: 'newanimal timmy bird' or 'query timmy speak'")
			continue
		}

		switch animalRequest[0] {
		case "newanimal":
			name, animalType := animalRequest[1], animalRequest[2]
			newAnimal, err := animalFactory(animalType)
			if err != nil {
				fmt.Println(err)
				continue
			}
			animals[name] = newAnimal
			fmt.Println("Created it!")
		case "query":
			animal, request := animalRequest[1], animalRequest[2]
			Request(animals[animal], request)
		default:
			fmt.Println("Only 'newanimal' or 'query' can be handled. Try again.")
		}

	}

}
