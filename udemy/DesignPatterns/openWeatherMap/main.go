package main

import (
	"fmt"
	"log"

	w "github.com/roca/GO/udemy/DesignPatterns/iopenWeatherMap/facade"
)

var MyAPIkey string = "62a3ce28cadac204dcbf12c78c41d4ee"

func main() {
	weatherMap := w.CurrentWeatherData{APIkey: MyAPIkey}
	weather, err := weatherMap.GetByCityAndCountryCode("Madrid", "ES")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Temperature in Madrid is %f celsius\n", weather.Main.Temp-273.15)
}
