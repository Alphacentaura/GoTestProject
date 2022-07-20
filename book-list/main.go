package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {

	router := mux.NewRouter()

	books = append(books, Book{ID: 1, Title: "Book 1", Author: "Author 1", Year: "2011"},
		Book{ID: 2, Title: "Book 2", Author: "Author 2", Year: "2012"},
		Book{ID: 3, Title: "Book 3", Author: "Author 3", Year: "2013"},
		Book{ID: 4, Title: "Book 4", Author: "Author 4", Year: "2017"},
		Book{ID: 5, Title: "Book 5", Author: "Author 5", Year: "2021"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatalln(http.ListenAndServe(":8000", router))

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println(params)

	for _, book := range books {
		bookId, _ := strconv.Atoi(params["id"])
		if book.ID == bookId {
			_ = json.NewEncoder(w).Encode(book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	log.Println(book)

	books = append(books, book)

	_ = json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}

	json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])

	for index, book := range books {
		if book.ID == i {
			books = append(books[:index], books[index+1:]...)
		}
	}

	json.NewEncoder(w).Encode(books)
}
