package prototype

import (
	"errors"
	"fmt"
)

type IShirtCloner interface {
	GetClone(m int) (IItemInfoGetter, error)
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

type ShirtCache struct{}

func (s *ShirtCache) GetClone(m int) (IItemInfoGetter, error) {
	switch m {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}

type IItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

func (sc *ShirtColor) String() string {
	lookup := map[ShirtColor]string{
		White: "White",
		Black: "Black",
		Blue:  "Blue",
	}
	return lookup[*sc]
}

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Price: %f, SKU: %s, Color: %s", s.Price, s.SKU, s.Color.String())
}

func GetShirtsCloner() IShirtCloner {
	return new(ShirtCache)
}

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}
var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}
var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

func (i *Shirt) GetPrice() float32 {
	return i.Price
}
