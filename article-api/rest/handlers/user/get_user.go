package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Age   int     `json:"age"`
	Phone *string `json:"phone"`
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []User

	err := h.dbConf.Select(&users, "SELECT id, name, email, age, phone FROM users")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
