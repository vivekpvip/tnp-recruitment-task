package handlers

import (
	"encoding/json"
	"event-backend/db"
	"event-backend/models"
	"fmt"
	"net/http"
)

// 📌 Register User
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO users (name, email, college, year, expectations) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = db.DB.QueryRow(query, user.Name, user.Email, user.College, user.Year, user.Expectations).Scan(&user.ID)
	if err != nil {
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "✅ Registration Successful!")
}

// 📌 Check-in User
func CheckInUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	query := `UPDATE users SET checked_in = TRUE WHERE email = $1`
	res, err := db.DB.Exec(query, user.Email)
	if err != nil {
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "User not found or already checked in", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "✅ Checked In!")
}

// 📌 Check-out User
func CheckOutUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	query := `UPDATE users SET checked_in = FALSE WHERE email = $1`
	res, err := db.DB.Exec(query, user.Email)
	if err != nil {
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "User not found or already checked out", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "✅ Checked Out!")
}
