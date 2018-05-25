//************************************************************************//
// API "cellar": Application Media Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/GOCODE/goa/design
// --out=$(GOPATH)/src/github.com/GOCODE/goa
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"unicode/utf8"
)

// bottle media type (default view)
//
// Identifier: application/vnd.gophercon.goa.bottle; view=default
type Bottle struct {
	// Unique bottle ID
	ID int `form:"ID" json:"ID" xml:"ID"`
	// Name of bottle
	Name string `form:"name" json:"name" xml:"name"`
	// Name of bottle
	Rating int `form:"rating" json:"rating" xml:"rating"`
	// Vintage of bottle
	Vintage int `form:"vintage" json:"vintage" xml:"vintage"`
}

// Validate validates the Bottle media type instance.
func (mt *Bottle) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	if utf8.RuneCountInString(mt.Name) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 2, true))
	}
	if mt.Rating < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, mt.Rating, 1, true))
	}
	if mt.Rating > 5 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, mt.Rating, 5, false))
	}
	if mt.Vintage < 1900 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.vintage`, mt.Vintage, 1900, true))
	}
	return
}

// DecodeBottle decodes the Bottle instance encoded in resp body.
func (c *Client) DecodeBottle(resp *http.Response) (*Bottle, error) {
	var decoded Bottle
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
