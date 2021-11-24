package main

import "fmt"

// OCP: Open/Closed Principle
// open for extension, closed for modification
// Specification pattern

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Filter struct {
	//
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product,0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product,0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	result := make([]*Product,0)
	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

type Product struct {
	name  string
	color Color
	size  Size
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c *ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s *SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	fmt.Printf("Green products (old):\n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}
}
