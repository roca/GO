// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Package toy contains support for managing toy inventory.
package toy

// Toy represents a toy we sell.
type Toy struct {
	Name   string
	Weight int

	onHand int
	sold   int
}

// New creates values of type toy.
func New(name string, weight int) *Toy {
	return &Toy{
		Name:   name,
		Weight: weight,
	}
}

// OnHand returns the current number of this
// toy on hand.
func (t *Toy) OnHand() int {
	return t.onHand
}

// UpdateOnHand updates the on hand count and
// returns the current value.
func (t *Toy) UpdateOnHand(count int) int {
	t.onHand = t.onHand + count
	return t.onHand
}

// Sold returns the current number of this
// toy sold.
func (t *Toy) Sold() int {
	return t.sold
}

// UpdateSold updates the sold count and
// returns the current value.
func (t *Toy) UpdateSold(count int) int {
	t.sold = t.sold + count
	return t.sold
}
