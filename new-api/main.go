package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world", 200)
	}))
	mux.Handle("GET /products", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 3)
		fmt.Fprintln(w, "products request", 200)
	}))
	mux.Handle("GET /users", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "user request", 200)
	}))
	fmt.Println("Server is listing..")
	http.ListenAndServe(":8000", mux)
}
