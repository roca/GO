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
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

type BetterFilter struct { }

func (bf *BetterFilter) FilterBy(products []Product, specifications ...Specification) []*Product {
	result := make([]*Product, 0)
	for i := range products {
		specIsSatisfiedCount := 0
		for _, spec := range specifications {
			if spec.IsSatisfied(&products[i]) {
				specIsSatisfiedCount++
			}
		}
		if specIsSatisfiedCount == len(specifications) {
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

type AndSpecification struct {
	first,second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	f := Filter{}

	// Old way
	fmt.Printf("Green and Small products (old):\n")
	for _, v := range f.FilterBySizeAndColor(products, small, green) {
		fmt.Printf(" - %s is green and small\n", v.name)
	}

	// New way
	// Specification pattern
	// BetterFilter struct is closed for modification but open for extension
	bf := BetterFilter{}
	fmt.Printf("Green and Small products (new):\n")
	greenSpec := ColorSpecification{green}
	sizeSpec := SizeSpecification{small}
	for _, v := range bf.FilterBy(products, &greenSpec, &sizeSpec) {
		fmt.Printf(" - %s is green and small\n", v.name)
	}

	fmt.Printf("Green and Small products (newnew):\n")
	andSpec := AndSpecification{&greenSpec, &sizeSpec}
	for _, v := range bf.FilterBy(products, &andSpec) {
		fmt.Printf(" - %s is green and small\n", v.name)
	}
}
