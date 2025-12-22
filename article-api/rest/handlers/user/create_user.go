package user

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating a user goes here
	user := struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Age   int    `json:"age"`
		Phone string `json:"phone"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Basic validation
	if user.Name == "" || user.Email == "" {
		http.Error(w, "Name and Email are required", http.StatusBadRequest)
		return
	}

	// SQL insert query
	query := `
		INSERT INTO users (name, email, age, phone)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	// Execute query
	err = h.dbConf.QueryRow(
		query,
		user.Name,
		user.Email,
		user.Age,
		user.Phone,
	).Scan(&user.ID)
	json.NewEncoder(w).Encode(user)
}
