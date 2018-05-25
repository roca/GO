package model

import (
	"encoding/json"
	"net/http"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Vendor struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Contact *Contact `json:"contact"`
}

func GetVendors() []*Vendor {
	var result []*Vendor

	port := os.Getenv("VENDOR_PORT")

	resp, err := http.Get("http://localhost:" + port + " /api/vendors?type=manufacturing")

	if err != nil {
		return result
	}

	data := make([]byte, resp.ContentLength)
	resp.Body.Read(data)

	json.Unmarshal(data, &result)

	return result
}
