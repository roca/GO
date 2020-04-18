package main

import (
	"fmt"
	"log"

	"example.com/openWeatherMap/facade"
)

var MyAPIkey string = "62a3ce28cadac204dcbf12c78c41d4ee"

func main() {
	weatherMap := facade.CurrentWeatherData{APIkey: MyAPIkey}
	weather, err := weatherMap.GetByCityAndCountryCode("New York", "US")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Temperature in New York is %f Fahrenheit\n", ((weather.Main.Temp-273.15)*(9.0/5.0))+32.0)
}
