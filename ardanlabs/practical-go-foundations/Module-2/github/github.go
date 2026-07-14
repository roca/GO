package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Given a github user login, return name and number of public repos
// curl -i https://api.github.com/users/ardanlabs
// curl -i -H 'User-Agent: go' https://api.github.com/users/ardanlabs

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	Name, NumRepos, err := UserInfo(ctx, "ardanlabs")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Name:", Name, ", NumRepos:", NumRepos)
}

func demo() {
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

}

// UserInfo return name and number of public repos from GitHub API.
func UserInfo(ctx context.Context, login string) (string, int, error) {
	url := "https://api.github.com/users/" + login

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	// resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%q - bad status: %s", url, resp.Request.Response.Status)
	}

	return parseResponse(resp.Body)
}

func parseResponse(r io.Reader) (string, int, error) {
	var reply struct {
		Name     string
		NumRepos int `json:"public_repos"`
	}
	dec := json.NewDecoder(r)
	if err := dec.Decode(&reply); err != nil {
		return "", 0, err
	}

	return reply.Name, reply.NumRepos, nil
}

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
