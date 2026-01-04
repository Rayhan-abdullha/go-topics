package middleware

import (
	"fmt"
	"net/http"
)

func AuthMid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("This is auth mid...")
		next.ServeHTTP(w, r)
	})
}
