package main

import (
	"event-backend/db"
	"event-backend/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db.ConnectDB() // Connect to PostgreSQL

	http.HandleFunc("/register", routes.RegisterHandler)
	http.HandleFunc("/checkin", routes.CheckInHandler)
	http.HandleFunc("/checkout", routes.CheckOutHandler)

	fmt.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
