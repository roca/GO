package composition

import "fmt"

type Athlete struct{}

func (a *Athlete) Train() {
	println("Training")
}

func Swim() {
	println("Swimming!")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    *func()
}

//--------------------------------------------------------

type ITrainer interface {
	Train()
}

type ISwimmer interface {
	Swim()
}

type SwimmerImplementor struct{}

func (s *SwimmerImplementor) Swim() {
	println("Swimming!")
}

type CompositeSwimmerB struct {
	ITrainer
	ISwimmer
}

//---------------------------------------------

type Animal struct{}

func (r *Animal) Eat() {
	println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}

//---------------------------------

type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}

//-------------------------------------
type Parent struct {
	SomeField int
}

type Son struct {
	P Parent
}

func GetParentField(p *Parent) int {
	fmt.Println(p.SomeField)
	return p.SomeField
}

// cannot use son (type Son) as type Parent in argument to GetParentField
