package routes

import (
	"database/sql"
	"net/http"

	"github.com/salwakhairu/Tugas-Besar-PBW-Sistem-Karyawan/controller"
	"github.com/salwakhairu/Tugas-Besar-PBW-Sistem-Karyawan/handlers"
)

// MapRoutes maps all HTTP routes
func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", handlers.LoginPage)
	server.HandleFunc("/signup", handlers.SignUpPage)
	server.HandleFunc("/signup/submit", handlers.SignUpHandler(db))
	server.HandleFunc("/check-username", handlers.CheckUsernameHandler(db))
	server.HandleFunc("/login", handlers.LoginPage)
	server.HandleFunc("/login/submit", handlers.LoginHandler(db))
	server.HandleFunc("/LandingPage", controller.DashboardPage)
	server.HandleFunc("/LandingPage/product", controller.NewIndexLandingProduct(db))
	server.HandleFunc("/cart", controller.ViewCart(db))
	server.HandleFunc("/cart/create", controller.NewCreateCartController(db))
	server.HandleFunc("/cart/update", controller.NewUpdateCartController(db))
	server.HandleFunc("/cart/delete", controller.NewDeleteCartController(db))
	server.HandleFunc("/checkout/success", controller.CheckoutSuccess(db))
	server.HandleFunc("/product", controller.NewIndexProduct(db))
	server.HandleFunc("/product/create", controller.NewCreateProductController(db))
	server.HandleFunc("/product/update", controller.NewUpdateProductController(db))
	server.HandleFunc("/product/delete", controller.NewDeleteProductController(db))
}
