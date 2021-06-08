package main

// Imports
import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Structs
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// TODO : Implement DB
var books []Book

/*
	Function that returns all the books
*/
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

/*
	Function that returns the book by its id
*/
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

// TODO
func createBook(w http.ResponseWriter, r *http.Request) {

}

// TODO
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// TODO
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

// Main func
func main() {
	router := mux.NewRouter()

	// Mock data
	books = append(books, Book{ID: "1", Isbn: "12345678", Title: "Book One", Author: &Author{ID: "1", FirstName: "Meša", LastName: "Selimović"}})
	books = append(books, Book{ID: "2", Isbn: "87654321", Title: "Book Two", Author: &Author{ID: "1", FirstName: "Meša", LastName: "Selimović"}})
	books = append(books, Book{ID: "3", Isbn: "45678321", Title: "Book Three", Author: &Author{ID: "2", FirstName: "Ivo", LastName: "Andrić"}})

	// Handlers
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
