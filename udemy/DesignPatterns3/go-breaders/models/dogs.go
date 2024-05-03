package models

type DogBreed struct {
	ID int `json:"id"`
	BreadProps
}

type Dog struct{
	ID int `json:"id"`
	DogName string `json:"dog_name"`
	BreedID int `json:"breed_id"`
	BreederID int `json:"breeder_id"`
	Color string `json:"color"`
}
