package api

import (
	"encoding/json"
	"log"
)

type Assignment struct {
	Page         string             `json:"page"`
	Words        []string           `json:"words"`
	Percentages  map[string]float64 `json:"percentages"`
	Special      []interface{}      `json:"special"`
	ExtraSpecial []interface{}      `json:"extra_special"`
}

func (a Assignment) GetResponse() string {
	asBytes, err := json.MarshalIndent(a, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	return string(asBytes)
}
