// Package athletes ... (your comments to describe the exported package)
package athletes

import (
	"strings"
)

// Info ... (your comments to describe the exported struct)
type Info struct {
	Country   string
	HairColor string
}

// Player ... (your comments to describe the exported struct)
type Player struct {
	Name  string
	Sport string
	Age   int
	Info
}

// ToLowercase ... (your comments to describe the exported method receiver)
func (p *Player) ToLowercase() *Player {
	p.Name = strings.ToLower(p.Name)
	p.Sport = strings.ToLower(p.Sport)
	p.Country = strings.ToLower(p.Country)
	p.HairColor = strings.ToLower(p.HairColor)

	return p
}
