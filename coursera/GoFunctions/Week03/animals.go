package main

import "fmt"

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

func main() {

}
