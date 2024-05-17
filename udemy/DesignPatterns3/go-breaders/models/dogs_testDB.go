package models

import "log"

func (m *testRepository) AllDogBreeds() ([]*DogBreed, error) {
	log.Println("AllDogBreeds() called")
	return []*DogBreed{}, nil
}
