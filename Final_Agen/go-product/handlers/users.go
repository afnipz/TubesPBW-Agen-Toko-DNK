package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path/filepath"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

// SignUpPage menampilkan halaman sign up
func SignUpPage(w http.ResponseWriter, r *http.Request) {
	fp := filepath.Join("views", "signup.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, "Unable to load sign up page", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Unable to render sign up page", http.StatusInternalServerError)
		return
	}
}

// SignUpHandler memproses form sign up
func SignUpHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Invalid form submission", http.StatusBadRequest)
				return
			}

			username := r.FormValue("username")
			password := r.FormValue("password")

			// Check if username already exists
			var exists bool
			err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM pengguna WHERE username = ?)", username).Scan(&exists)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			if exists {
				http.Error(w, "Username sudah digunakan", http.StatusConflict)
				return
			}

			// Hash password sebelum disimpan di database
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			// Simpan username dan hashed password di database
			_, err = db.Exec("INSERT INTO pengguna (username, password) VALUES (?, ?)", username, hashedPassword)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			// Arahkan pengguna ke halaman login setelah berhasil mendaftar
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// CheckUsernameHandler untuk memeriksa apakah username sudah ada
func CheckUsernameHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		var exists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM pengguna WHERE username = ?)", username).Scan(&exists)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		response := map[string]bool{"exists": exists}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

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

			// Simpan username dalam cookie
			http.SetCookie(w, &http.Cookie{
				Name:  "username",
				Value: username,
				Path:  "/",
			})

			http.Redirect(w, r, "/LandingPage/product", http.StatusSeeOther)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
