package main

import "net/http"

// Given a github user login, return name and number of public repos
// curl -i https://api.github.com/users/ardanlabs

func main() {
	resp, err := http.Get("https://api.github.com/users/ardanlabs")
}
