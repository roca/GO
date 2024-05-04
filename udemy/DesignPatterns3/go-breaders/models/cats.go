package models

type CatBreed struct {
	ID int `json:"id"`
	BreadProps
}

type Cat struct {
	ID int `json:"id"`
	PetProps
	Breed CatBreed `json:"breed"`
}
