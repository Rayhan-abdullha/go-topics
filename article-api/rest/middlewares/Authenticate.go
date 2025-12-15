package middlewares

import (
	"fmt"
	"net/http"
)

// import "net/http"

func (m *Middleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Auth mid")
		next.ServeHTTP(w, r)
	})
}
