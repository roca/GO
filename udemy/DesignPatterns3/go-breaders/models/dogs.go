package models

type DogBreed struct {
	ID int `json:"id"`
	BreadProps
}

type Dog struct {
	ID int `json:"id"`
	PetProps
	Breed DogBreed `json:"breed"`
}

func (d *DogBreed) All() ([]*DogBreed, error) {
	return d.AllDogBreeds()
}
