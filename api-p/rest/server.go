package rest

import (
	"fmt"
	"net/http"
	"server/config"
	"server/rest/handler"
	"server/rest/handler/product"
	mid "server/rest/middleware"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
	// userHandler    *user.User
}

func NewServer(
	cnf *config.Config,
	productHandler *product.Handler,
	// userHandler *user.User
) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
	}
}

func (server *Server) Start() {
	mux := http.NewServeMux()
	mng := mid.NewManager()
	mng.Use(
		mid.PreflightReq,
		mid.Cors,
		mid.LoggerMid,
	)

	mux.Handle("POST /register", mng.With(
		http.HandlerFunc(handler.Registration),
	))
	mux.Handle("GET /profile", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqQuery := r.URL.Query()
		name := reqQuery.Get("name")
		email := reqQuery.Get("email")
		password := reqQuery.Get("password")
		if name == "" || email == "" || password == "" {
			fmt.Fprintln(w, "credential error", 400)
			return
		}
		if email == "r@gmail.com" && name == "rayhan" && password == "1234" {
			fmt.Fprintln(w, "This is the profile page.")
		} else {
			fmt.Fprintln(w, "You are unathorize", 401)
		}
	}))
	server.productHandler.RegisterRoutes(mux, mng)

	muxWrapper := mng.WrapMux(mux)
	fmt.Println("Server is running on port 4000")
	http.ListenAndServe(":4000", muxWrapper)
}
