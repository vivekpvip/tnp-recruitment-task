package routes

import (
	"event-backend/handlers"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		handlers.RegisterUser(w, r)
	}
}

func CheckInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		handlers.CheckInUser(w, r)
	}
}

func CheckOutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		handlers.CheckOutUser(w, r)
	}
}
