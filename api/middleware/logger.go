package middleware

import (
	"fmt"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	fmt.Println("logger... execute")
	return next
}
