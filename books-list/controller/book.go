package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mo-mirzania/api/books-list/model"
	bookRepository "github.com/mo-mirzania/api/books-list/repository/book"
	"github.com/mo-mirzania/api/books-list/utils"
)

// Controller struct
type Controller struct {
}

var books []model.Book

// GetBooks func
func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book model.Book
		var error model.Error
		books = []model.Book{}
		bookRepo := bookRepository.BookRepository{}
		books, err := bookRepo.GetBooks(db, book, books)
		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Add("Content-type", "application/json")
		utils.SendSuccess(w, books)
	}
}

// GetBook func
func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book model.Book
		var error model.Error
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			json.NewEncoder(w).Encode("Invalid ID")
			return
		}
		bookRepo := bookRepository.BookRepository{}
		book, err = bookRepo.GetBook(db, book, id)
		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Not found"
				utils.SendError(w, http.StatusNotFound, error)
				return
			} else {
				error.Message = "Server error"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		}
		w.Header().Add("Content-type", "application/json")
		utils.SendSuccess(w, book)
	}
}

// AddBook func
func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book model.Book
		var error model.Error
		bookRepo := bookRepository.BookRepository{}
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		book, err = bookRepo.AddBook(db, book)
		if err != nil {
			error.Message = "Sever error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Add("Content-type", "application/json")
		utils.SendSuccess(w, book)
	}
}

// UpdateBook func
func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book model.Book
		var error model.Error
		bookRepo := bookRepository.BookRepository{}
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		book, err = bookRepo.UpdateBook(db, book)
		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Add("Content-type", "application/json")
		utils.SendSuccess(w, book)
	}
}

// RemoveBook func
func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var error model.Error
		bookRepo := bookRepository.BookRepository{}
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result, err := bookRepo.RemoveBook(db, id)
		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Add("Content-type", "application/json")
		utils.SendSuccess(w, result)
	}
}
