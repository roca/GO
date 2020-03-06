package composition

type Athlete struct{}

func (a *Athlete) Train() {
	println("Training")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    *func()
}

//--------------------------------------------------------

type Trainer interface {
	Train()
}

type Swimmer interface {
	Swim()
}

type SwimmerImplementor struct{}

func (s *SwimmerImplementor) Swim() {
	println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
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
