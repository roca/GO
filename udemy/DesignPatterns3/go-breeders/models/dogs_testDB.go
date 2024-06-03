package models

import "log"

func (m *testRepository) AllDogBreeds() ([]*DogBreed, error) {
	log.Println("AllDogBreeds() called")
	return []*DogBreed{}, nil
}

func (m *testRepository) GetBreedByName(b string) (*DogBreed, error) {
	log.Println("GetBreedByName() called")
	return &DogBreed{}, nil
}
