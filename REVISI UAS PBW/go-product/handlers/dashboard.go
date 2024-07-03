package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Product struct {
	Id        string
	Nama      string
	Harga     string
	Deskripsi string
}

// DashboardPage menampilkan halaman dashboard dengan produk unggulan
func DashboardPage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, nama, harga, deskripsi FROM product")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var products []Product
		for rows.Next() {
			var product Product
			err = rows.Scan(&product.Id, &product.Nama, &product.Harga, &product.Deskripsi)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			products = append(products, product)
		}

		fp := filepath.Join("views", "dashboard.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			Products []Product
		}{
			Products: products,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
