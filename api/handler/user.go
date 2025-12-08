package handler

import (
	"article/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func CorsAndPrifix(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.Header().Set("Content-Type", "application/json")
}
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	CorsAndPrifix(w, r)
	if r.Method == "OPTIONS" {
		fmt.Println(r.Method)
		w.WriteHeader(http.StatusOK)
		return
	}

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

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	CorsAndPrifix(w, r)
	if db.DB == nil {
		db.DB = append(db.DB, db.User{
			Name:  "rayhan",
			Email: "rayhan@gmail.com",
		})
	}
	json.NewEncoder(w).Encode(db.GetUser())
}
