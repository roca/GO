// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/bYY-TRjfH0

// Sample program to show how functions can return multiple values while using
// named and struct types.
package main

import (
	"encoding/json"
	"fmt"
)

// user is a struct type that declares user information.
type user struct {
	ID   int
	Name string
}

// main is the entry point for the application.
func main() {
	// Retrieve the user profile.
	u, err := retrieveUser("sally")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display the user profile.
	fmt.Printf("%+v\n", *u)
}

// retrieveUser retrieves the user document for the specified
// user and returns a pointer to a user type value.
func retrieveUser(name string) (*user, error) {
	// Make a call to get the user in a json response.
	r, err := getUser(name)
	if err != nil {
		return nil, err
	}

	// Unmarshal the json document into a value of
	// the user struct type.
	var u user
	err = json.Unmarshal([]byte(r), &u)
	return &u, err
}

// GetUser simulates a web call that returns a json
// document for the specified user.
func getUser(name string) (string, error) {
	response := `{"id":1432, "name":"sally"}`
	return response, nil
}
