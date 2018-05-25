//************************************************************************//
// API "cellar": Application User Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/GOCODE/goa/design
// --out=$(GOPATH)/src/github.com/GOCODE/goa
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// BottlePayLoad is the type used to create bottles
type bottlePayLoad struct {
	// Name of bottle
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Name of bottle
	Rating *int `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
	// Vintage of bottle
	Vintage *int `form:"vintage,omitempty" json:"vintage,omitempty" xml:"vintage,omitempty"`
}

// Validate validates the bottlePayLoad type instance.
func (ut *bottlePayLoad) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Vintage == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "vintage"))
	}
	if ut.Rating == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "rating"))
	}

	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 2, true))
		}
	}
	if ut.Rating != nil {
		if *ut.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *ut.Rating, 1, true))
		}
	}
	if ut.Rating != nil {
		if *ut.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *ut.Rating, 5, false))
		}
	}
	if ut.Vintage != nil {
		if *ut.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.vintage`, *ut.Vintage, 1900, true))
		}
	}
	return
}

// Publicize creates BottlePayLoad from bottlePayLoad
func (ut *bottlePayLoad) Publicize() *BottlePayLoad {
	var pub BottlePayLoad
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Rating != nil {
		pub.Rating = *ut.Rating
	}
	if ut.Vintage != nil {
		pub.Vintage = *ut.Vintage
	}
	return &pub
}

// BottlePayLoad is the type used to create bottles
type BottlePayLoad struct {
	// Name of bottle
	Name string `form:"name" json:"name" xml:"name"`
	// Name of bottle
	Rating int `form:"rating" json:"rating" xml:"rating"`
	// Vintage of bottle
	Vintage int `form:"vintage" json:"vintage" xml:"vintage"`
}

// Validate validates the BottlePayLoad type instance.
func (ut *BottlePayLoad) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	if utf8.RuneCountInString(ut.Name) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, ut.Name, utf8.RuneCountInString(ut.Name), 2, true))
	}
	if ut.Rating < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, ut.Rating, 1, true))
	}
	if ut.Rating > 5 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, ut.Rating, 5, false))
	}
	if ut.Vintage < 1900 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.vintage`, ut.Vintage, 1900, true))
	}
	return
}
