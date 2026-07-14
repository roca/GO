package main

import "fmt"

type IProductInfoRetriever interface {
	GetPrice() float32
	GetName() string
}

type IVisitor interface {
	Visit(IProductInfoRetriever)
}

type IVisitable interface {
	Accept(IVisitor)
}

type Product struct {
	Price float32
	Name  string
}

func (p *Product) GetPrice() float32 {
	return p.Price
}
func (p *Product) GetName() string {
	return p.Name
}
func (p *Product) Accept(v IVisitor) {
	v.Visit(p)
}

type Rice struct {
	Product
}
type Pasta struct {
	Product
}

type PriceVisitor struct {
	Sum float32
}

func (pv *PriceVisitor) Visit(p IProductInfoRetriever) {
	pv.Sum += p.GetPrice()
}

type NameVisitor struct {
	ProductList string
}

func (n *NameVisitor) Visit(p IProductInfoRetriever) {
	n.ProductList = fmt.Sprintf("%s\n%s", p.GetName(), n.ProductList)
}

type Fridge struct{
	Product
}
func (f *Fridge) GetPrice() float32 {
	return f.Product.Price + 20
}
func (f *Fridge) Accept(v IVisitor){
	v.Visit(f)
}

func main() {
	products := make([]IVisitable, 3)
	products[0] = &Rice{Product: Product{Price: 32.0, Name: "Some Rice"}}
	products[1] = &Pasta{Product: Product{Price: 40.0, Name: "Some Pasta"}}
	products[2] = &Fridge{Product: Product{Price: 50.0, Name: "A Fridge"}}

	priceVisitor := &PriceVisitor{}
	for _, p := range products {
		p.Accept(priceVisitor)
	}
	fmt.Printf("\nTotal: %f", priceVisitor.Sum)

	nameVisitor := &NameVisitor{}
	for _, p := range products {
		p.Accept(nameVisitor)
	}
	fmt.Printf("\nProduct list:\n--------------------\n%s", nameVisitor.ProductList)

}
