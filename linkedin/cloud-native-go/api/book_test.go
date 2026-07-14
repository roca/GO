package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"}
	bytes := book.ToJSON()
	assert.Equal(t, `{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"0123456789"}`, string(bytes), "Book JSON marshalling wrong")
}

func TestBookFromJSON(t *testing.T) {
	
	bytes := []byte(`{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"0123456789"}`)
	book := FromJSON(bytes)
	expectedBook := Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"}

	assert.Equal(t, expectedBook, book, "Book titles did not match")
}
