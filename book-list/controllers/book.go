package controllers

import (
	"GoTest/book-list/driver"
	"GoTest/book-list/models"
	BookRepository "GoTest/book-list/repository/book"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Controller struct{}

var books []models.Book

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		books = []models.Book{}
		bookRepo := BookRepository.BookRepository{}
		books = bookRepo.GetBooks(db, book, books)

		json.NewEncoder(w).Encode(books)
	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		params := mux.Vars(r)

		id, err := strconv.Atoi(params["id"])
		driver.LogFatal(err)

		bookRepo := BookRepository.BookRepository{}
		book = bookRepo.GetBook(db, book, id)

		json.NewEncoder(w).Encode(book)
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var bookId int

		json.NewDecoder(r.Body).Decode(&book)

		bookRepo := BookRepository.BookRepository{}
		bookId = bookRepo.AddBook(db, book)

		json.NewEncoder(w).Encode(bookId)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		json.NewDecoder(r.Body).Decode(&book)

		bookRepo := BookRepository.BookRepository{}
		rowsUpdated := bookRepo.UpdateBook(db, book)

		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		id, err := strconv.Atoi(params["id"])
		driver.LogFatal(err)

		bookRepo := BookRepository.BookRepository{}
		rowsDeleted := bookRepo.RemoveBook(db, id)

		json.NewEncoder(w).Encode(rowsDeleted)
	}
}
