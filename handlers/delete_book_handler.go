package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteBook(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("Content-Type", "application-json")

	params := mux.Vars(r)

	result, err := db.Exec("DELETE FROM books WHERE id = ?", params["id"])

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

	fmt.Fprintln(w, "Row deleted successfully", noOfRowsAffected)

}
