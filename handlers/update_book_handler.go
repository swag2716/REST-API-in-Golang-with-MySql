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

func UpdateBook(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application-json")

	params := mux.Vars(r)

	var updateBook models.Book

	err := json.NewDecoder(r.Body).Decode(&updateBook)

	if err != nil {
		log.Fatal(err)
	}

	result, err := db.Exec("UPDATE books SET title = ?, author = ? WHERE id = ?", updateBook.Title, updateBook.Author, params["id"])

	if err != nil {
		log.Fatal(err)
	}

	noOfRowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	if noOfRowsAffected == 0 {
		fmt.Fprintln(w, "No book exist with id :", params["id"])
		return
	}

	fmt.Fprintln(w, "Book updated at id :", params["id"])

}
