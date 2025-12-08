package main

import (
	"article/handler"
	"fmt"
	"log"
	"net/http"
	"time"
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		next.ServeHTTP(w, r)
		log.Println(time.Since(t))
	})
}
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a := r.Header.Get("a")
		if a == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		fmt.Fprintln(w, "Authentication success")
		next.ServeHTTP(w, r)
	})
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world")
	}))
	// mux.Handle("GET /users", http.HandlerFunc(handler.GetUserHandler))
	// add middleare
	mux.Handle("GET /users", middleware(
		AuthMiddleware(http.HandlerFunc(handler.GetUserHandler)),
	))

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
