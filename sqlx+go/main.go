package main

import (
	"article/db"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	mux := http.NewServeMux()
	db, err := db.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Database is connected..")

	mux.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello server", 200)
	}))
	mux.Handle("GET /users", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := []User{}
		err := db.Select(&user, "SELECT id, name, email, age FROM users")
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Faild to fetch student", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}))

	mux.Handle("POST /users", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := User{}
		encode := json.NewDecoder(r.Body)
		err := encode.Decode(&user)
		if err != nil {
			http.Error(w, "Invalid request", 401)
			return
		}
		query := `
			INSERT INTO users (name, email, age)-
			VALUES($1, $2, $3)
			RETURNING id
		`
		err = db.QueryRow(
			query,
			user.Name,
			user.Email,
			user.Age,
		).Scan(&user.ID)
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(user)
	}))
	fmt.Println("Server is listing port 4001")
	http.ListenAndServe(":4001", mux)
}
