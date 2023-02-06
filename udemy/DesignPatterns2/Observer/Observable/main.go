package main

import "container/list"

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

type Person struct {
	Observable
	Name string
}

func (p *Person) CatchACold() {
	p.Notify(p.Name + " has a cold")
}

type DoctorService struct{
}

func (d *DoctorService) Notify(data interface{}) {
	println("A doctor has been called to see", data.(string))
}

func NewPerson(name string) *Person {
	return &Person{
		Observable: Observable{subs: list.New()},
		Name:       name,
	}
}

func main() {
	p := NewPerson("Boris")
	ds := &DoctorService{}
	p.Subscribe(ds)

	p.CatchACold()
}
