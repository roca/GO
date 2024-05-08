package main

import "fmt"

// Animal is the type for our abstract factory
type Animal interface {
	Says()
	LikesWater() bool
}

// Dog is the concrete factory for dogs
type Dog struct{}

// implement the abstract factory for dog
func (d *Dog) Says() {
	fmt.Println("Woof")
}
func (d *Dog) LikesWater() bool {
	return true
}

// Cat is the concrete factory for cats
type Cat struct{}

// implement the abstract factory for cat
func (c *Cat) Says() {
	fmt.Println("Meow")
}

func (c *Cat) LikesWater() bool {
	return false
}

type AnimalFactory interface {
	New() Animal
}

type DogFatcory struct{}

func (df *DogFatcory) New() Animal {
	return &Dog{}
}

type CatFatcory struct{}

func (cf *CatFatcory) New() Animal {
	return &Cat{}
}

func main() {
	var dogFactory, catFactory AnimalFactory
	dogFactory = &DogFatcory{}
	catFactory = &CatFatcory{}

	dog := dogFactory.New()
	cat := catFactory.New()

	dog.Says()
	fmt.Println("A dog like water:",dog.LikesWater())

	cat.Says()
	fmt.Println("A cat like water:", cat.LikesWater())

}
