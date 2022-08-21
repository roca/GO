package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetRateLimit() (string, error) {
	response, err := http.Get("http://localhost:8080/ratelimit")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d: %s",response.StatusCode,string(body)), nil
}

func main() {
	f := func(i int) {
		body, err := GetRateLimit()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d: %s\n", i, body)
	}
	for i := 0; i < 1000; i++ {
		f(i)
		time.Sleep(200 * time.Millisecond)
	}

}
