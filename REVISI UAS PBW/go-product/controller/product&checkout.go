package controller

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

type Product struct {
	Id        string
	Nama      string
	Harga     string
	Deskripsi string
}

func NewIndexProduct(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
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

			err = rows.Scan(
				&product.Id,
				&product.Nama,
				&product.Harga,
				&product.Deskripsi,
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			products = append(products, product)
		}

		fp := filepath.Join("views", "index.html")
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

func NewUpdateProductController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			id := r.URL.Query().Get("id")
			r.ParseForm()

			nama := r.Form["nama"][0]
			harga := r.Form["harga"][0]
			deskripsi := r.Form["deskripsi"][0]
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
			err := row.Scan(
				&product.Nama,
				&product.Harga,
				&product.Deskripsi,
			)
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

// DashboardPage menampilkan halaman dashboard
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

			err = rows.Scan(
				&product.Id,
				&product.Nama,
				&product.Harga,
				&product.Deskripsi,
			)
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

type CartItem struct {
	ID          string
	ProductName string
	Kuantitas   int
	Produk      Product
	Total       int
}

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

			http.Redirect(w, r, "/LandingPage/product#Unggulan", http.StatusSeeOther)
			return
		} else if r.Method == "GET" {
			rows, err := db.Query("SELECT id, product_name, kuantitas FROM cart")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer rows.Close()

			var cartItems []CartItem
			for rows.Next() {
				var cart CartItem
				err := rows.Scan(&cart.ID, &cart.ProductName, &cart.Kuantitas)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				cartItems = append(cartItems, cart)
			}

			fp := filepath.Join("views", "cartcreate.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = tmpl.Execute(w, cartItems)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func ViewCart(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`
            SELECT c.id, c.product_name, c.kuantitas, p.nama, p.harga, p.deskripsi 
            FROM cart c 
            JOIN product p ON c.product_name = p.nama`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var cartItems []CartItem
		for rows.Next() {
			var item CartItem
			err := rows.Scan(&item.ID, &item.ProductName, &item.Kuantitas, &item.Produk.Nama, &item.Produk.Harga, &item.Produk.Deskripsi)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Konversi harga produk ke tipe int
			harga, err := strconv.Atoi(item.Produk.Harga)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			item.Total = item.Kuantitas * harga // Hitung total harga
			cartItems = append(cartItems, item)
		}

		fp := filepath.Join("views", "cart.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			CartItems []CartItem
		}{
			CartItems: cartItems,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func NewUpdateCartController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			id := r.FormValue("id")
			kuantitasStr := r.FormValue("kuantitas")
			kuantitas, err := strconv.Atoi(kuantitasStr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			_, err = db.Exec("UPDATE cart SET kuantitas=? WHERE id=?", kuantitas, id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/cart", http.StatusSeeOther)
			return
		} else if r.Method == "GET" {
			id := r.URL.Query().Get("id")

			row := db.QueryRow(`
                SELECT c.id, c.product_name, c.kuantitas, p.nama, p.harga, p.deskripsi 
                FROM cart c 
                JOIN product p ON c.product_name = p.nama 
                WHERE c.id = ?`, id)
			if row.Err() != nil {
				http.Error(w, row.Err().Error(), http.StatusInternalServerError)
				return
			}

			var item CartItem
			err := row.Scan(&item.ID, &item.ProductName, &item.Kuantitas, &item.Produk.Nama, &item.Produk.Harga, &item.Produk.Deskripsi)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fp := filepath.Join("views", "cartupdate.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			data := struct {
				Cart CartItem
			}{
				Cart: item,
			}

			err = tmpl.Execute(w, data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func NewDeleteCartController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "ID parameter is required", http.StatusBadRequest)
			return
		}

		_, err := db.Exec("DELETE FROM cart WHERE id = ?", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/cart", http.StatusSeeOther)
	}
}

type CheckoutData struct {
	CartItems []CartItem
	Total     int
	Timestamp string
	Username  string
}

func CheckoutSuccess(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ambil username dari cookie
		cookie, err := r.Cookie("username")
		if err != nil {
			http.Error(w, "Unable to retrieve username", http.StatusInternalServerError)
			return
		}
		username := cookie.Value
		fmt.Println("Username:", username) // Tambahkan ini untuk debugging

		// Retrieve cart items from the database
		rows, err := db.Query(`
            SELECT c.id, c.product_name, c.kuantitas, p.harga 
            FROM cart c 
            JOIN product p ON c.product_name = p.nama
        `)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var cartItems []CartItem
		var totalAmount int

		// Prepare the SQL transaction
		tx, err := db.Begin()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer func() {
			if err != nil {
				tx.Rollback()
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Commit the transaction if no errors occurred
			err = tx.Commit()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Clear the cart after successful checkout
			_, err = db.Exec("DELETE FROM cart")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}()

		// Iterate through cart items and calculate totals
		for rows.Next() {
			var cartItem CartItem
			var product Product

			// Scan data from row
			err := rows.Scan(&cartItem.ID, &cartItem.ProductName, &cartItem.Kuantitas, &product.Harga)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Assign related product data
			cartItem.Produk = product

			// Convert product price to int
			harga, err := strconv.Atoi(product.Harga)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			cartItem.Total = cartItem.Kuantitas * harga // Calculate total price
			cartItems = append(cartItems, cartItem)
			totalAmount += cartItem.Total
		}

		// Format timestamp to "YYYY-MM-DD HH:mm:ss"
		timestamp := time.Now().Format("2006-01-02 15:04:05")

		// Data for template
		data := CheckoutData{
			CartItems: cartItems,
			Total:     totalAmount,
			Timestamp: timestamp,
			Username:  username,
		}

		// Render template
		fp := filepath.Join("views", "checkout_success.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// DashboardPage menampilkan halaman dashboard
func DashboardPageAdmin(w http.ResponseWriter, r *http.Request) {
	fp := filepath.Join("views", "LandingPageAdmin.html")
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
