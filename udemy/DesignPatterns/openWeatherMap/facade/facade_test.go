package openWeatherMap

import (
	"bytes"
	
	"io"
	"testing"
)

func getMockData() io.Reader {
	response := `{
		"coord": {
		  "lon": -122.08,
		  "lat": 37.39
		},
		"weather": [
		  {
			"id": 800,
			"main": "Clear",
			"description": "clear sky",
			"icon": "01d"
		  }
		],
		"base": "stations",
		"main": {
		  "temp": 282.55,
		  "feels_like": 281.86,
		  "temp_min": 280.37,
		  "temp_max": 284.26,
		  "pressure": 1023,
		  "humidity": 100
		},
		"visibility": 16093,
		"wind": {
		  "speed": 1.5,
		  "deg": 350
		},
		"clouds": {
		  "all": 1
		},
		"dt": 1560350645,
		"sys": {
		  "type": 1,
		  "id": 5122,
		  "message": 0.0139,
		  "country": "US",
		  "sunrise": 1560343627,
		  "sunset": 1560396563
		},
		"timezone": -25200,
		"id": 420006353,
		"name": "Mountain View",
		"cod": 200
	  }
	  `

	return bytes.NewReader([]byte(response))
}

func TestOpenWeatherMap_responseParser(t *testing.T) {
	r := getMockData()
	openWeatherMap := CurrentWeatherData{APIkey: ""}
	weather, err := openWeatherMap.responseParser(r)
	if err != nil {
		t.Fatal(err)
	}
	if weather.ID != 420006353 {
		t.Errorf("This city id is 420006353, not %d\n", weather.ID)
	}

}

