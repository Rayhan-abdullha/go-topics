package main

import (
	"article/handler"
	"article/middleware"
	"article/utils"
	"fmt"
	"net/http"
)

type Middlewares func(http.Handler) http.Handler

func Middleware(next http.Handler, mid ...Middlewares) http.Handler {
	for _, m := range mid {
		next = m(next)
	}
	return next
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world")
	}))
	// mux.Handle("GET /users", http.HandlerFunc(handler.GetUserHandler))
	// add middleare
	mux.Handle("GET /users", Middleware(
		utils.Handler(handler.GetUserHandler),
		middleware.Logger,
		middleware.AuthMiddleware,
	))

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
