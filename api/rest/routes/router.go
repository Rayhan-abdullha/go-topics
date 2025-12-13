package routes

import (
	handler "article/rest/handlers"
	middleware "article/rest/middlewares"
	"article/utils"
	"fmt"
	"net/http"
)

func InitialRoute(mux *http.ServeMux, m *middleware.MiddleWareManager) {
	mux.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world")
	}))
	// mux.Handle("GET /users", http.HandlerFunc(handler.GetUserHandler))
	// add middleare
	mux.Handle("GET /users", m.With(
		utils.Handler(handler.GetUserHandler),
	))
	mux.Handle("POST /users", m.With(
		http.HandlerFunc(handler.CreateUserHandler),
	))
}
