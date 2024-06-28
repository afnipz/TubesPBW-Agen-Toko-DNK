package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() *sql.DB {
	dsn := "root:@tcp(localhost:3306)/agen"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to open database connection:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	return db
}
