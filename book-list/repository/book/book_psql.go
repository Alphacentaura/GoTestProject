package BookRepository

import (
	"GoTest/book-list/driver"
	"GoTest/book-list/models"
	"database/sql"
)

type BookRepository struct{}

func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) []models.Book {
	rows, err := db.Query("select * from books")
	driver.LogFatal(err)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		driver.LogFatal(err)
	}(rows)

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		driver.LogFatal(err)

		books = append(books, book)
	}
	return books
}

func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) models.Book {
	rows := db.QueryRow("select * from books where id=$1", id)

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	driver.LogFatal(err)

	return book
}

func (b BookRepository) AddBook(db *sql.DB, book models.Book) int {
	var bookId int
	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id", book.Title, book.Author, book.Year).Scan(&bookId)
	driver.LogFatal(err)
	return bookId
}

func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) int64 {
	result, err := db.Exec("update books set title=$1, author = $2, year = $3 where id = $4 RETURNING id", &book.Title, &book.Author, &book.Year, &book.ID)

	rowsUpdated, err := result.RowsAffected()
	driver.LogFatal(err)

	return rowsUpdated
}

func (b BookRepository) RemoveBook(db *sql.DB, id int) int64 {
	result, err := db.Exec("delete from books where id = $1", id)
	driver.LogFatal(err)

	rowsDeleted, err := result.RowsAffected()
	driver.LogFatal(err)

	return rowsDeleted
}
