package book

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

//BooksHandleFunc ... serves json for book store
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		allBooks := AllBooks()
		writeJSON(w, allBooks)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		book := FromJSON(body)
		isbn, created := CreateBook(book)
		if created {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request type"))
	}
	// b, err := json.Marshal(books)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	// w.Header().Add("Content-Type", "application/json; charset=utf-8")
	// w.Write(b)
}

func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/books/"):]
	log.Printf(isbn)
	switch method := r.Method; method {
	case http.MethodGet:
		book, found := GetBook(isbn)
		if !found {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusFound)
			writeJSON(w, book)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		exists := UpdateBook(isbn, book)
		if exists {
			w.WriteHeader(http.StatusOK)
			book, _ = GetBook(isbn)
			writeJSON(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		DeleteBook(isbn)
		log.Print(books)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

// GetBook ... returns the book for a given ISBN
func GetBook(isbn string) (Book, bool) {
	book, found := books[isbn]
	log.Println(book)
	return book, found
}

func AllBooks() []Book {
	var allBooks []Book
	for _, b := range books {
		allBooks = append(allBooks, b)
	}
	return allBooks
}

// CreateBook creates a new Book if it does not exist
func CreateBook(book Book) (string, bool) {
	_, exists := books[book.ISBN]
	if exists {
		return "", false
	}
	books[book.ISBN] = book
	return book.ISBN, true
}

// UpdateBook updates an existing book
func UpdateBook(isbn string, book Book) bool {
	_, exists := books[isbn]
	if exists {
		books[isbn] = book
	}
	return exists
}

// DeleteBook removes a book from the map by ISBN key
func DeleteBook(isbn string) {
	delete(books, isbn)
}
