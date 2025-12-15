package rest

import (
	"article/config"
	"article/rest/handlers/user"
	middleware "article/rest/middlewares"
	"fmt"
	"net/http"
)

type Server struct {
	userHandler *user.Handlers
}

func NewServer(
	userHandler *user.Handlers,
) *Server {
	return &Server{
		userHandler: userHandler,
	}
}

func (server *Server) Start(conf config.Config) {
	// global middleware
	m := middleware.NewManager()
	m.Use(
		middleware.CorsAndPrefix,
		middleware.Logger,
		middleware.AuthMiddleware,
	)

	mux := http.NewServeMux()

	h := m.WrapMux(mux)

	// user route
	server.userHandler.RegisterRoutes(mux, m)

	port := ":" + conf.HttpPort
	fmt.Println("Listening on port ", port)
	http.ListenAndServe(port, h)
}
