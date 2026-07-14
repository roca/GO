package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/*
from : // https://data-asg.goldprice.org/dbxRates/USD
{
	ts: 1682067394727,
	tsj: 1682067385985,
	date: "Apr 21st 2023, 04:56:25 am NY",
	items: [
		{
		curr: "USD",
		xauPrice: 1988.0825,
		xagPrice: 25.1213,
		chgXau: -17.1825,
		chgXag: -0.1332,
		pcXau: -0.8569,
		pcXag: -0.5274,
		xauClose: 2005.265,
		xagClose: 25.2545
		}
	]
}
*/

var currency = "USD"

type Gold struct {
	Prices []Price `json:"items"`
	Client *http.Client
}

type Price struct {
	Currency      string    `json:"curr"`
	Price         float64   `json:"xauPrice"`
	Change        float64   `json:"chgXau"`
	PreviousClose float64   `json:"xauClose"`
	Time          time.Time `json:"-"`
}

func (g *Gold) GetPrices() (*Price, error) {
	if g.Client == nil {
		g.Client = &http.Client{}
	}

	client := g.Client

	url := fmt.Sprintf("https://data-asg.goldprice.org/dbxRates/%s", currency)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("error getting gold prices: %v", err)
		return nil, fmt.Errorf("error getting gold prices: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Reading body of the response: %v", err)
		return nil, fmt.Errorf("error getting gold prices: %v", err)
	}

	gold := Gold{}
	var previous, current, change float64
	err = json.Unmarshal(body, &gold)
	if err != nil {
		log.Printf("error unmarshalling gold prices: %v", err)
		return nil, fmt.Errorf("error getting gold prices: %v", err)
	}

	previous, current, change = gold.Prices[0].PreviousClose, gold.Prices[0].Price, gold.Prices[0].Change

	return &Price{
		Currency:      currency,
		Price:         current,
		Change:        change,
		PreviousClose: previous,
		Time:          time.Now(),
	}, nil

}
