package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Client struct {
	FirstName string `required:"true" json:"first_name"`
	LastName  string `json:"last_name"`
}

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Client *Client `json:"client"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for idx, book := range books {
		if book.ID == params["id"] {
			books = append(books[:idx], books[idx+1:]...)
			json.NewEncoder(w).Encode(book.ID)
			break
		}
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = uuid.NewString()
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for idx, book := range books {
		if book.ID == params["id"] {
			books = append(books[:idx], books[idx+1:]...)
			var b Book
			_ = json.NewDecoder(r.Body).Decode(&b)
			books = append(books, b)
			json.NewEncoder(w).Encode(b)
			break
		}
	}
}

func main() {
	r := mux.NewRouter()

	books = append(books, Book{
		ID:    uuid.NewString(),
		Isbn:  "543534",
		Title: "First Book",
		Client: &Client{
			FirstName: "John",
			LastName:  "Doe",
		},
	})

	books = append(books, Book{
		ID:    uuid.NewString(),
		Isbn:  "65634",
		Title: "Second Book",
		Client: &Client{
			FirstName: "Smith",
			LastName:  "Fallen",
		},
	})

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Printf("Starting server at http://localhost:8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
