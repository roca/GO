package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"google.com",
		"facebook.com",
		"stackoverflow.com",
		"golang.org",
		"amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(newLink string) {
			time.Sleep(5 * time.Second)
			checkLink(newLink, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get("http://" + link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
