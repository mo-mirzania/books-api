package bookRepository

import (
	"database/sql"
	"fmt"

	"github.com/mo-mirzania/api/books-list/model"
)

// BookRepository struct
type BookRepository struct{}

// GetBooks method
func (b BookRepository) GetBooks(db *sql.DB, book model.Book, books []model.Book) ([]model.Book, error) {
	rows, err := db.Query("select * from books;")
	if err != nil {
		return []model.Book{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		books = append(books, book)
	}

	if err != nil {
		return []model.Book{}, err
	}
	return books, nil
}

// GetBook func
func (b BookRepository) GetBook(db *sql.DB, book model.Book, id int) (model.Book, error) {
	row := db.QueryRow(fmt.Sprintf("select * from books where id = %d;", id))
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

// AddBook func
func (b BookRepository) AddBook(db *sql.DB, book model.Book) (model.Book, error) {
	insertQuery := `INSERT INTO books (id, title, author, year) VALUES ($1, $2, $3, $4);`
	_, err := db.Exec(insertQuery, book.ID, book.Title, book.Author, book.Year)
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

// UpdateBook method
func (b BookRepository) UpdateBook(db *sql.DB, book model.Book) (model.Book, error) {
	insertQuery := `UPDATE books set title = $1, author = $2, year = $3 WHERE id = $4;`
	_, err := db.Exec(insertQuery, book.Title, book.Author, book.Year, book.ID)
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

// RemoveBook method
func (b BookRepository) RemoveBook(db *sql.DB, id int) (int64, error) {
	deleteQuery := `DELETE FROM books WHERE id = $1;`
	result, err := db.Exec(deleteQuery, id)
	if err != nil {
		return 0, err
	}
	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsDeleted, nil
}
