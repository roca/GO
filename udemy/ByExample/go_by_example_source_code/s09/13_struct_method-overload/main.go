// File name: ...\s09\13_struct_method-overload\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"strings"
)

type movie struct {
	name  string
	actor string
}

type imdb struct {
	movie
	name    string
	comment string
}

func (m movie) fullInfo() string {
	return m.name + "-" + m.actor
}

func (i imdb) fullInfo() string {
	return strings.ToLower(i.movie.name + "-" + 
		i.movie.actor + "-" + i.name + "-" + i.comment)
}

func main() {
	m := movie{"Forrest Gump", "Tom Hanks"}
	i1 := imdb{m, "Sony Pictures", "Loved it!"}

	i2 := imdb{
		movie: movie{
			name:  "The Godfather",
			actor: "Marlon Brando",
		},
		name:    "Paramount Pictures",
		comment: "Superb!",
	}

	fmt.Println(m.fullInfo())
	fmt.Println(i1.fullInfo())
	fmt.Println(i2.fullInfo())

	fmt.Printf("%s-%s \n", i2.movie.name, i2.name)

	p1 := &movie{"Forrest Gump", "Tom Hanks"}
	fmt.Printf("%T - %v - name=%s - actor=%s\n", p1, *p1, p1.name, p1.actor)

	p2 := &i2
	fmt.Printf("%T - %v\n", p2, *p2)
}
