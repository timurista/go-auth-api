package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native go", Author: "M L Reimer", ISBN: "12345"}
	json := book.ToJSON()
	assert.Equal(t, `{"title":"Cloud Native go","author":"M L Reimer","isbn":"12345"}`, string(json), "Book JSON marshalling wrong")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"Cloud Native go","author":"M L Reimer","isbn":"12345"}`)
	book := FromJSON(json)
	assert.Equal(t, Book{Title: "Cloud Native go", Author: "M L Reimer", ISBN: "12345"}, book, "Book JSON unmarshalling wrong")
}
