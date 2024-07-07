package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

// type Reply struct {
// 	Name            string `json:"name"`
// 	PublicRepoCount int    `json:"public_repos"`
// }

func main() {
	ctx, canmcel := context.WithTimeout(context.Background(), 3*time.Millisecond)
	defer canmcel()

	name, count, err := githubInfo(ctx, "roca")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Name:", name)
	fmt.Println("PublicRepoCount:", count)


	fmt.Println(githubInfo(ctx, "tebeka"))
}

// githubInfo returns the name and public repo count of a GitHub user.
func githubInfo(ctx context.Context, login string) (string, int, error) {
	url := "https://api.github.com/users/" + url.PathEscape(login)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", 0, err
	}

	// resp, err := http.Get(fmt.Sprintf(url))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%#v - %s", url, resp.Status)
	}

	var r struct { // anonymous struct
		Name            string `json:"name"`
		PublicRepoCount int    `json:"public_repos"`
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}

	return r.Name, r.PublicRepoCount, nil
}

/* JSON <-> Go
true/false <-> true/false
string <-> string
null <-> nil
number <-> float64, float32, int8, int16, int32, int64, int, uint8, ...
array <-> []any ([]interface{})
object <-> map[string]any, struct

encoding/json API
JSON -> io.Reader -> Go: json.Decoder
JSON -> []byte -> Go: json.Unmarshal
Go -> io.Writer -> JSON: json.Encoder
Go -> []byte -> JSON: json.Marshal
*/
