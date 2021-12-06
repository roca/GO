package main

import "fmt"

/*
   DIP: Dependency Inversion Principle:

   DIP is a software design principle that states that
   "High-level modules should not depend on low-level modules.
   Both should depend on abstractions."
*/

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// low-level module

type RelationshipBrowser interface {
   FindAllChildrenOf(name string) []*Person
}

type Relationships struct {
	relations []Info
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
   var children []*Person

   for i, rel := range r.relations {
		if rel.relationship == Parent && rel.from.name == name {
			children = append(children, r.relations[i].to)
		}
	}

   return children
}

func (r *Relationships) AddParentAndChild(parent *Person, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// high-level module
type Research struct {
	// breaks the DIP principle
	// relationships *Relationships
   browser RelationshipBrowser
}

func (r *Research) Investigate() { // find all children of a parent
	for _, p := range r.browser.FindAllChildrenOf("John") {
      fmt.Println("John has a child called ", p.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{&relationships}
	r.Investigate()
}
