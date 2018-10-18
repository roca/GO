package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"google.com",
		"facebok.com",
		"stackoverflow.com",
		"golang.org",
		"amazon.com",
	}

	for _, link := range links {
		go checkLink(link)
	}

}

func checkLink(link string) {
	_, err := http.Get("http://" + link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}
