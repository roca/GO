package composition

import "testing"


func TestAthlete_Train(t *testing.T) {
	athlete := Athlete{}
	athlete.Train()
}

func TestSwimmer_Swim(t *testing.T) {
	localSwim := Swim
	swimmer := CompositeSwimmerA{
		MySwim: &localSwim,
	}
	swimmer.MyAthlete.Train()
	(*swimmer.MySwim)()
}

func TestAnimal_Swim(t *testing.T) {
	fish := Shark{
		Swim: Swim,
	}
	fish.Eat()
	fish.Swim()
}

func TestSwimmer_Swim2(t *testing.T) {
	swimmer := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImplementor{},
	}
	swimmer.Train()
	swimmer.Swim()
}
