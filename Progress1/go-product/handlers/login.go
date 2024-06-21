package handlers

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

// LoginPage menampilkan halaman login
func LoginPage(w http.ResponseWriter, r *http.Request) {
	fp := filepath.Join("views", "login.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, "Unable to load login page", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Unable to render login page", http.StatusInternalServerError)
		return
	}
}

// LoginHandler memproses form login
func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Invalid form submission", http.StatusBadRequest)
				return
			}

			username := r.FormValue("username")
			password := r.FormValue("password")

			var hashedPassword string
			err = db.QueryRow("SELECT password FROM pengguna WHERE username = ?", username).Scan(&hashedPassword)
			if err != nil {
				if err == sql.ErrNoRows {
					http.Error(w, "Invalid username or password", http.StatusUnauthorized)
					return
				}
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
			if err != nil {
				http.Error(w, "Invalid username or password", http.StatusUnauthorized)
				return
			}

			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
