package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggerMid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		next.ServeHTTP(w, r)
		end := time.Since(t)
		log.Println(end)
	})
}
