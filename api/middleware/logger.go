package middleware

import (
	"article/utils"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return utils.Handler(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		next.ServeHTTP(w, r)
		log.Println(time.Since(t))
		fmt.Fprintln(w, time.Since(t))
	})
}
