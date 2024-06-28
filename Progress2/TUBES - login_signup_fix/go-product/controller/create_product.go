package controller

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

func NewCreateProductController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()

			name := r.Form["nama"][0]
			harga := r.Form["harga"][0]
			deskripsi := r.Form["deskripsi"][0]
			_, err := db.Exec("INSERT INTO product (nama, harga, deskripsi) VALUES (?, ?, ?)", name, harga, deskripsi)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/product", http.StatusMovedPermanently)
			return
		} else if r.Method == "GET" {
			rows, err := db.Query("SELECT id, nama, harga, deskripsi FROM product")
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer rows.Close()

			var products []Product
			for rows.Next() {
				var product Product
				err := rows.Scan(&product.Id, &product.Nama, &product.Harga, &product.Deskripsi)
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				products = append(products, product)
			}

			fp := filepath.Join("views", "create.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = tmpl.Execute(w, products)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
