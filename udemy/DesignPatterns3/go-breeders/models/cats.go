package models

type CatBreed struct {
	ID int `json:"id" xml:"id"`
	BreadProps
}

type Cat struct {
	ID int `json:"id" xml:"id"`
	PetProps
	Breed CatBreed `json:"breed" xml:"breed"`
}
