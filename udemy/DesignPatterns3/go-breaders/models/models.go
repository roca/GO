package models

import "time"

type Breeder struct{}

type Pet struct{}

type BreadProps struct {
	Breed            string `json:"breed"`
	WeightLowLbs     int    `json:"weight_low_lbs"`
	WeightHighLbs    int    `json:"weight_high_lbs"`
	AverageWeightLbs int    `json:"average_weight_lbs"`
	LifeSpan         int    `json:"life_span"`
	Details          string `json:"details"`
	AlternateNames   string `json:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin"`
}
type PetProps struct {
	Name             string    `json:"name"`
	BreedID          int       `json:"breed_id"`
	BreederID        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DataOfBirth      time.Time `json:"data_of_birth"`
	SpayedOrNeutered int       `json:"spayed_or_neutered"`
	Description      string    `json:"description"`
	Breeder          Breeder   `json:"breeder"`
}
