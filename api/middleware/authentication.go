package middleware

import (
	"fmt"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// before request
		fmt.Println("Before handler...")

		if true {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Unauthorized")
			return
		} else {
			next.ServeHTTP(w, r) // call next handler

			// after request
			fmt.Println("After handler...")
		}
	})
}
