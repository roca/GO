package models

import (
	"database/sql"
	"time"
)

var repo Repository

type Models struct {
	DogBreed DogBreed
}

func New(conn *sql.DB) *Models {
	if conn != nil {
		repo = newMysqlRepository(conn)
	} else {
		repo = newTestRepository(conn)
	}

	return &Models{
		DogBreed: DogBreed{},
	}
}

type Breeder struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Address     string      `json:"address"`
	City        string      `json:"city"`
	ProvState   string      `json:"prov_state"`
	Country     string      `json:"country"`
	Postcode    string      `json:"postcode"`
	PhoneNumber string      `json:"phone_number"`
	Email       string      `json:"email"`
	Active      int         `json:"active"`
	DogBreeds   []*DogBreed `json:"dog_breeds"`
	CatBreeds   []*CatBreed `json:"cat_breeds"`
}

type Pet struct {
	Species     string `json:"species"`
	Breed       string `json:"breed"`
	MaxWeight   int    `json:"max_weight"`
	MinWeight   int    `json:"min_weight"`
	Description string `json:"description"`
	LifeSpan    int    `json:"life_span"`
}

type BreadProps struct {
	Breed            string `json:"breed" xml:"breed"`
	WeightLowLbs     int    `json:"weight_low_lbs" xml:"weight_low_lbs"`
	WeightHighLbs    int    `json:"weight_high_lbs" xml:"weight_high_lbs"`
	AverageWeightLbs int    `json:"average_weight_lbs" xml:"average_weight_lbs"`
	LifeSpan         int    `json:"life_span" xml:"life_span"`
	Details          string `json:"details" xml:"details"`
	AlternateNames   string `json:"alternate_names" xml:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin" xml:"geographic_origin"`
}
type PetProps struct {
	Name             string    `json:"name"`
	BreedID          int       `json:"breed_id"`
	BreederID        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DataOfBirth      time.Time `json:"data_of_birth"`
	SpayedOrNeutered int       `json:"spayed_or_neutered"`
	Description      string    `json:"description"`
	Weight           int       `json:"weight"`
	Breeder          Breeder   `json:"breeder"`
}
