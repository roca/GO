package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Given a github user login, return name and number of public repos
// curl -i https://api.github.com/users/ardanlabs
// curl -i -H 'User-Agent: go' https://api.github.com/users/ardanlabs

func main() {
	resp, err := http.Get("https://api.github.com/users/ardanlabs")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("ERROR: bad status - %s\n", resp.Status)
		os.Exit(1)
	}

	ctype := resp.Header.Get("Content-Type")
	fmt.Println("content-type:", ctype)

	// io.Copy(os.Stdout, resp.Body)
	var reply struct {
		Name     string
		NumRepos int `json:"public_repos"`
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&reply); err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	fmt.Println(reply.Name, reply.NumRepos)

	/* JSON <-> Go

	Types
	string <-> string
	true/false <-> bool
	number <-> float[64,32....], int[8,64......], uint[8....]
	array <-> []T, []any
	object <-> map[string]any, struct

	encoding/json API
	JSON -> []byte -> Go: Unmarshal
	Go -> []byte -> JSON: Marshal

	JSON -> io.Reader -> Go: Decoder
	Go -> io.Writer -> JSON Encoder
	*/

}
