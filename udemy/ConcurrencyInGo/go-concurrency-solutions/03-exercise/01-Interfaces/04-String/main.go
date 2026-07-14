package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) String() string {
	return fmt.Sprintf("%s <%s>", u.name, u.email)
}

// TODO: Implement custom formating for user struct values.

func main() {
	var s fmt.Stringer
	u := user{
		name:  "John Doe",
		email: "johndoe@example.com",
	}
	s = u
	fmt.Println(s)
}
