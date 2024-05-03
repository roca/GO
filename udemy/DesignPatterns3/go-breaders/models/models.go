package models

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
