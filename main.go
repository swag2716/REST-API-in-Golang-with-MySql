package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swapnika/Accessing-MySql-in-Golang-through-router/database"
	"github.com/swapnika/Accessing-MySql-in-Golang-through-router/handlers"
)

var db *sql.DB

func main() {
	db = database.InitDB()
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetBooks(w, r, db)
	}).Methods("GET")
	router.HandleFunc("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetBook(w, r, db)
	}).Methods("GET")
	router.HandleFunc("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateBook(w, r, db)
	}).Methods("PATCH")
	router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateBook(w, r, db)
	}).Methods("POST")
	router.HandleFunc("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteBook(w, r, db)
	}).Methods("DELETE")

	fmt.Println("Running on server 8000")
	http.ListenAndServe(":8000", router)

}
