package main

import "fmt"

// LSP: Liskov Substitution Principle
// LSP is a programming principle that states that objects in a program
// should be replaceable with instances of their subtypes without altering the correctness of that program.

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width  int
	height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}
func (r *Rectangle) SetWidth(width int) {
	r.width = width
}
func (r *Rectangle) GetHeight() int {
	return r.height
}
func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.height = height
	s.width = height
}

type Square2 struct {
	size int // width and height
}
func (s *Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}


func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := width * 10
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected an area of ", expectedArea, ", but got ", actualArea, "\n")
}

func main() {
	rc := &Rectangle{2,3}
	UseIt(rc)

	sq2 := Square2{5}
	r := sq2.Rectangle()
	UseIt(&r)
}
