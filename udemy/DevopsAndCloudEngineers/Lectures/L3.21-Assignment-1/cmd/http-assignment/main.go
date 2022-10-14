package main

import (
	"encoding/json"
	"fmt"

	"github.com/roca/GO/tree/staging/udemy/DevopsAndCloudEngineers/Lectures/L3.21-Assignment-1/pkg/api"
)

func main() {
	a := api.Assignment{
		Page: "assignment1",
		Words: []string{ "eigth", "two", "two", "two", "five"},
		Special: []interface{}{ "one", "two", nil},
		ExtraSpecial: []interface{}{ 1, 2, "3"},
		Percentages: map[string]float64{"eigth": 0.5, "five": 1, "two": 0.66},
	}

	asBytes, err := json.MarshalIndent(a, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(asBytes))

}
