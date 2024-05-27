package pets

import "go-breaders/models"

func New(species string) *models.Pet {
	return &models.Pet{
		Species:     species,
		Breed:       "",
		MinWeight:   0,
		MaxWeight:   0,
		Description: "no description entered yet",
		LifeSpan:    0,
	}
}
