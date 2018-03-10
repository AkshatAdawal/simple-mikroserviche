package book

import (
	"encoding/json"
	"net/http"
)

//Book ... Canonical representation of book
type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty	"`
}

//ToJSON ... returns json representation of books
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		return nil
	}
	return ToJSON
}

//FromJSON ... Creates a book from json string
func FromJSON(b []byte) Book {
	book := Book{}
	err := json.Unmarshal(b, &book)
	if err != nil {
		panic(err)
	}
	return book
}

var books = map[string]Book{
	"Three": Book{"One", "Two", "Three", "This starts a count"},
	"Six":   Book{"Four", "Five", "Six", "This starts from previous one"},
}

//BooksHandleFunc ... serves json for book store
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

func BookHandleFunc(w http.ResponseWriter, r *http.Request) {

}
