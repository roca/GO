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

// Cow : type called Cow which is a struct containing three fields:food, locomotion, and noise, all of which are strings
type Cow struct {
	food       string
	locomotion string
	noise      string
}

// Eat method should print the Cow’s food
func (a Cow) Eat() {
	fmt.Println(a.food)
}

// Move method should print the Cow’s locomotion
func (a Cow) Move() {
	fmt.Println(a.locomotion)
}

// Speak method should print the Cow’s spoken sound
func (a Cow) Speak() {
	fmt.Println(a.noise)
}

// Bird : type called Cow which is a struct containing three fields:food, locomotion, and noise, all of which are strings
type Bird struct {
	food       string
	locomotion string
	noise      string
}

// Eat method should print the Cow’s food
func (a Bird) Eat() {
	fmt.Println(a.food)
}

// Move method should print the Cow’s locomotion
func (a Bird) Move() {
	fmt.Println(a.locomotion)
}

// Speak method should print the Cow’s spoken sound
func (a Bird) Speak() {
	fmt.Println(a.noise)
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
		fmt.Println("This Cow doesn't know how to do that! It can only eat, move or speak")
	}
}

func animalFactory(animalType string) (Animal, error) {
	switch animalType {
	case "cow":
		return Cow{food: "grass", locomotion: "walk", noise: "moo"}, nil
	case "bird":
		return Bird{food: "worms", locomotion: "fly", noise: "peep"}, nil
	case "snake":
		return Cow{food: "mice", locomotion: "slither", noise: "hsss"}, nil
	}
	return nil, errors.New("Can't create animal of this type")
}

func main() {

	animals = map[string]Animal{}

	// Scanf wont work if your input has spaces :)
	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter two fields, the Cow of interest (cow, bird or snake) and the requested information (eat, move or speak) seprated by a space. (Example: cow eat)")

	for {
		fmt.Print(">")
		userInput, _ := consoleReader.ReadString('\n')
		userInput = strings.TrimSuffix(userInput, "\n")

		AnimalRequest := strings.Split(userInput, " ")

		animal, request := AnimalRequest[0], AnimalRequest[1]
		newAnimal, err := animalFactory(animal)
		if err != nil {
			fmt.Println(err)
		}
		animals[animal] = newAnimal

		Request(animals[animal], request)
	}

}
