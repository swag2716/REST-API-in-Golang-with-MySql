package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/swapnika/Accessing-MySql-in-Golang-through-router/models"
)

func CreateBook(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application-json")

	var newBook models.Book

	err := json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		log.Fatal(err)
	}

	result, err := db.Exec("INSERT INTO books (title, author) VALUES (?, ?)", newBook.Title, newBook.Author)

	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, "New book is inserted successfule at id :", id)
}
