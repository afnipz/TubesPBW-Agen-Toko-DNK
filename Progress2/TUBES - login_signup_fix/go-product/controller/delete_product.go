package controller

import (
	"database/sql"
	"net/http"
)

func NewDeleteProductController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get ID parameter from query string
		id := r.URL.Query().Get("id")

		// Execute delete query
		_, err := db.Exec("DELETE FROM product WHERE id = ?", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to product list page
		http.Redirect(w, r, "/product", http.StatusSeeOther)
	}
}
