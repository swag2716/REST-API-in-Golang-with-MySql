package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/swapnika/Accessing-MySql-in-Golang-through-router/models"
)

func GetBooks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")

	var books []models.Book

	rows, err := db.Query("SELECT * FROM books")

	if err != nil {
		fmt.Println("Here", err)
		log.Fatal(err)
	}

	for rows.Next() {
		var book models.Book

		err := rows.Scan(&book.Id, &book.Title, &book.Author)

		if err != nil {
			log.Fatal(err)
		}

		books = append(books, book)
	}

	if rows.Err() != nil {
		json.NewEncoder(w).Encode(models.Book{})
	}

	json.NewEncoder(w).Encode(books)

}
