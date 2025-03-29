package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := "user=postgres password=yourpassword dbname=eventdb sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("🚨 Database connection failed:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("🚨 Database not reachable:", err)
	}

	fmt.Println("✅ Connected to Database!")
}
