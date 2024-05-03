package models

type CatBreed struct {
	ID int `json:"id"`
	BreadProps
}

type Cat struct {
	ID        int    `json:"id"`
	CatName   string `json:"cat_name"`
	BreedID   int    `json:"breed_id"`
	BreederID int    `json:"breeder_id"`
	Color     string `json:"color"`
}
