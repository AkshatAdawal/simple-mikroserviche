package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"simple-mikroserviche/simple-mikroserviche/book"
)

// "simple-mikroserviche/simple-mikroserviche/hello"
func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		return ":8089"
	}
	fmt.Printf("Using port %v\n", port)
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, string(book.Book{Title: "Lolz", Author: "Lolz", ISBN: "Lolz"}.ToJSON()))
	log.Println("Served Page")
}

func echo(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query()["message"][0]
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, msg)
	log.Printf("Served echo API; Wrote %q", msg)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)
	http.HandleFunc("/api/books", book.BooksHandleFunc)
	http.HandleFunc("/api/books/", book.BookHandleFunc)

	http.ListenAndServe(port(), nil)
}
