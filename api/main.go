package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name  string
	Email string
}
type DB struct {
	Users []User
}

func (DB) CreateDB() DB {
	return DB{
		Users: []User{},
	}
}
func (d *DB) createUser(user User) {
	d.Users = append(d.Users, user)
}
func (d DB) getUser() []User {
	return d.Users
}

func main() {
	mux := http.NewServeMux()
	db := DB{}
	db.createUser(User{Name: "rayhan", Email: "rayhan@gmail.com"})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			json.NewEncoder(w).Encode(db.getUser())

		}
		if r.Method == "POST" {
			var user User
			dec := json.NewDecoder(r.Body)
			err := dec.Decode(&user)
			if err != nil || user.Name == "" || user.Email == "" {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}
			db.createUser(user)
			w.WriteHeader(201)
			json.NewEncoder(w).Encode(user)
		}
	})

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
