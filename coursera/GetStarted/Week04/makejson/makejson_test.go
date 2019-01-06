package main

import (
	"testing"
)

func TestReturnJSONString(t *testing.T) {
	myMap := map[string]string{
		"name":    "Romel Campbell",
		"address": "71 Seneca ave, Dumont NJ, 07628",
	}

	jsonString, _ := ReturnJSONString(myMap)

	expectedJSONString := []byte(`{"address":"71 Seneca ave, Dumont NJ, 07628","name":"Romel Campbell"}`)

	if jsonString != string(expectedJSONString) {
		t.Errorf("Could not covert this object to JSON string : %v", myMap)
		t.Error(jsonString, "\n")
		t.Error(string(expectedJSONString))
	}

}
