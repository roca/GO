package products

import "time"

type Product struct {
	ProductName string
	CreatedAt   time.Time
	UpdateAt    time.Time
}

// Factory function
func New(productName string) *Product {
	return &Product{
		ProductName: productName,
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}
}