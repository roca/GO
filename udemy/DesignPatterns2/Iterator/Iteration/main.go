package main

import "fmt"

type Person struct {
	FirstName, MiddleName, LastName string
}

// Case 1: Array

func (p *Person) Names() [3]string {
	return [3]string{p.FirstName, p.MiddleName, p.LastName}
}

// Case 2: Iterator

func (p *Person) NamesGenerator() <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		ch <- p.FirstName
		if len(p.MiddleName) > 0 {
			ch <- p.MiddleName
		}
		ch <- p.LastName
	}()
	return ch
}


// Case 3: Iterator

type PersonNamesIterator struct {
	person  *Person
	current int
}

func NewPersonNamesIterator(person *Person) *PersonNamesIterator {
	return &PersonNamesIterator{person, -1}
}

func (p *PersonNamesIterator) MoveNext() bool {
	p.current++
	return p.current < 3
}

func (p *PersonNamesIterator) Value() string {
	switch p.current {
	case 0:
		return p.person.FirstName
	case 1:
		return p.person.MiddleName
	case 2:
		return p.person.LastName
	}
	panic("out of bounds")
}

func main() {
	p := Person{"Alexander", "Graham", "Bell"}
	for it := NewPersonNamesIterator(&p); it.MoveNext(); {
		fmt.Println(it.Value())
	}
}
