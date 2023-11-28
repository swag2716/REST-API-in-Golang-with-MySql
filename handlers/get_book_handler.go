package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swapnika/Accessing-MySql-in-Golang-through-router/models"
)

func GetBook(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	row := db.QueryRow("SELECT * FROM books where id = ?", params["id"])

	var book models.Book

	err := row.Scan(&book.Id, &book.Title, &book.Author)

	if err != nil {
		if sql.ErrNoRows != nil {
			fmt.Fprintln(w, "No book exist with id :", params["id"])
			return
		}
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(book)

}
