package main

import (
	"container/list"
	"fmt"
)

// Observer and Observable Design Patterns

type Observable struct {
	subs *list.List
}

type Observer interface {
	Notify(data interface{})
}

type Subscriber struct {
}

func (s *Subscriber) Notify(data interface{}) {
	println(data.(string))
}

func (o *Observable) Subscribe(s Observer) {
	o.subs.PushBack(s)
}

func (o *Observable) Notify(data interface{}) {
	for e := o.subs.Front(); e != nil; e = e.Next() {
		e.Value.(Observer).Notify(data)
	}
}

func (o *Observable) Unsubscribe(s Observer) {
	for e := o.subs.Front(); e != nil; e = e.Next() {
		if e.Value.(Observer) == s {
			o.subs.Remove(e)
			break
		}
	}
}

type PropertyChange struct {
	Name  string
	Value interface{}
}

type Person struct {
	Observable
	age int
}

// Age() , SetAge()

func (p *Person) Age() int {
	return p.age
}

func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}
	oldCanVote := p.CanVote()
	p.age = age
	p.Notify(PropertyChange{"Age", age})
	if oldCanVote != p.CanVote() {
		p.Notify(PropertyChange{"CanVote", p.CanVote()})
	}
}
func (p *Person) CanVote() bool {
	return p.age >= 18
}

func NewPerson(age int) *Person {
	return &Person{
		Observable: Observable{subs: list.New()},
		age:        age,
	}
}

type TrafficManagement struct {
	o Observable
}

func (t *TrafficManagement) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok && pc.Name == "Age" {
		if pc.Value.(int) >= 16 {
			println("Congrats, you can drive now!")
			t.o.Unsubscribe(t)
		}
	}
}

// ElectoralRoll Observer
type ElectoralRoll struct {
	o Observable
}

func (e *ElectoralRoll) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok && pc.Name == "Age" {
		if pc.Value.(int) >= 18 {
			println("Congrats, you can vote now!")
			e.o.Unsubscribe(e)
		}
	}
}

func main() {
	p := NewPerson(0)
	er := &ElectoralRoll{p.Observable}
	p.Subscribe(er)

	for i := 1; i <= 20; i++ {
		fmt.Println("Setting age to ", i)
		p.SetAge(i)
	}

}
