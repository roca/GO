package models


type Dog struct {
	ID int `json:"id"`
	PetProps
	Breed DogBreed `json:"breed"`
}

func (d *Dog) GetDogOfMonthByID(id int) (*DogOfMonth, error) {
	return repo.GetDogOfMonthByID(id)
}

type DogOfMonth struct {
	ID int
	Dog *Dog
	Video string
	Image string
}

type DogBreed struct {
	ID int `json:"id"`
	BreadProps
}

func (d *DogBreed) All() ([]*DogBreed, error) {
	return repo.AllDogBreeds()
}

func (d *DogBreed) GetBreedByName(b string) (*DogBreed, error) {
	return repo.GetBreedByName(b)
}


