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
	return string(body), nil
}

func main() {
	f := func(i int) {
		body, err := GetRateLimit()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d: %s\n", i, body)
	}
	now := time.Now()
	fiveSeconds := time.Second * 5
	for i := 0; i < 10 && time.Since(now) < fiveSeconds; i++ {
		f(i)
	}
}
