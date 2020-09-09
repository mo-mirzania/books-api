package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/mo-mirzania/api/books-list/controller"
	drivers "github.com/mo-mirzania/api/books-list/driver"
	"github.com/mo-mirzania/api/books-list/model"
	"github.com/mo-mirzania/api/books-list/utils"
)

var books []model.Book

func main() {
	db := drivers.ConnectDB()
	controller := controller.Controller{}
	r := mux.NewRouter()
	r.HandleFunc("/ping", utils.Ping)
	r.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	r.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	r.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	r.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	r.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
