package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

// inisialisasi struct Product untuk merepresentasikan data produk
type Product struct {
	Id        string
	Nama      string
	Harga     string
	Deskripsi string
}

// Fungsi untuk menangani permintaan ke halaman index produk
func NewIndexProduct(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Melakukan query ke database untuk mengambil data produk
		rows, err := db.Query("SELECT id, nama, harga, deskripsi FROM product")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var products []Product
		// Memproses hasil query dan menyimpan ke dalam slice products
		for rows.Next() {
			var product Product
			err = rows.Scan(&product.Id, &product.Nama, &product.Harga, &product.Deskripsi)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			products = append(products, product)
		}

		// Menggabungkan path untuk template HTML
		fp := filepath.Join("views", "index.html")
		// Parsing file template HTML
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Membuat struktur data yang akan dikirim ke template
		data := struct {
			Products []Product
		}{
			Products: products,
		}

		// Mengeksekusi template dengan data produk
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// Fungsi untuk menangani permintaan pembaruan produk
func NewUpdateProductController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			id := r.URL.Query().Get("id")
			r.ParseForm()

			nama := r.Form["nama"][0]
			harga := r.Form["harga"][0]
			deskripsi := r.Form["deskripsi"][0]
			// Melakukan pembaruan data produk di database
			_, err := db.Exec("UPDATE product SET nama=?, harga=?, deskripsi=? WHERE id=?", nama, harga, deskripsi, id)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/product", http.StatusMovedPermanently)
			return
		} else if r.Method == "GET" {
			id := r.URL.Query().Get("id")

			row := db.QueryRow("SELECT nama, harga, deskripsi FROM product WHERE id = ?", id)
			if row.Err() != nil {
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var product Product
			err := row.Scan(&product.Nama, &product.Harga, &product.Deskripsi)
			product.Id = id
			if err != nil {
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			fp := filepath.Join("views", "update.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			data := make(map[string]any)
			data["product"] = product

			err = tmpl.Execute(w, data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}

// Fungsi untuk menangani permintaan penghapusan produk
func NewDeleteProductController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		_, err := db.Exec("DELETE FROM product WHERE id = ?", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/product", http.StatusSeeOther)
	}
}

// Fungsi untuk menangani permintaan pembuatan produk baru
func NewCreateProductController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()

			name := r.Form["nama"][0]
			harga := r.Form["harga"][0]
			deskripsi := r.Form["deskripsi"][0]
			// Menambahkan data produk baru ke database
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

// Fungsi untuk menampilkan halaman dashboard
func DashboardPage(w http.ResponseWriter, r *http.Request) {
	fp := filepath.Join("views", "dashboard.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Fungsi untuk menampilkan produk di halaman landing
func NewIndexLandingProduct(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
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

// Definisi struct CartItem untuk merepresentasikan item di keranjang belanja
type CartItem struct {
	ID          string
	ProductName string
	Kuantitas   int
	Produk      Product
	Total       int
}

// Fungsi untuk menangani permintaan pembuatan item keranjang baru
func NewCreateCartController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()

			productName := r.Form["ProductName"][0]
			kuantitas, err := strconv.Atoi(r.Form["Kuantitas"][0])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			_, err = db.Exec("INSERT INTO cart (product_name, kuantitas) VALUES (?, ?)", productName, kuantitas)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/cart", http.StatusMovedPermanently)
			return
		} else if r.Method == "GET" {
			fp := filepath.Join("views", "create_cart.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = tmpl.Execute(w, nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

// Fungsi untuk menampilkan halaman keranjang
func CartPage(w http.ResponseWriter, r *http.Request) {
	fp := filepath.Join("views", "cart.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
