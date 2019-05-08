package api

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	ISBN       string `json:"isbn"`
	Descrption string `json:"description,omitempty"`
}

func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

var books = []Book{
	Book{Title: "A good day to die", Author: "me", ISBN: "12345"},
	Book{Title: "Cloud Native Go", Author: "someone", ISBN: "12123"},
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}

	b, err := json.Marshal(books)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

// AllBooks returns a slice of all books
func AllBooks() []Book {
	values := make([]Book, len(books))
	idx := 0
	for _, book := range books {
		values[idx] = book
		idx++
	}
	return values
}

func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(books)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
