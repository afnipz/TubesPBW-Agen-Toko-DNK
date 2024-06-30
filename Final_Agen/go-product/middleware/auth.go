package middleware

import (
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Periksa apakah cookie username ada
		_, err := r.Cookie("username")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Lanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)
	})
}
