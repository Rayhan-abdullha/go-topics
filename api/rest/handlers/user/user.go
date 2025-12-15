package user

import (
	"article/db"
	"encoding/json"
	"net/http"
)

func (h *Handlers) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user db.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if user.Name == "" || user.Email == "" {
		http.Error(w, "Name & Email required", http.StatusBadRequest)
		return
	}

	db.CreateUser(user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *Handlers) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	if db.DB == nil {
		db.DB = append(db.DB, db.User{
			Name:  "rayhan",
			Email: "rayhan@gmail.com",
		})
	}

	users := db.GetUser()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
