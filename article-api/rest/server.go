package rest

import (
	"article/config"
	"article/rest/handlers/user"
	"article/rest/middlewares"
	"fmt"
	"net/http"
)

type Server struct {
	conf        *config.Config
	userHandler *user.Handler
}

func NewServer(conf *config.Config, userHandler *user.Handler) *Server {
	return &Server{
		conf:        conf,
		userHandler: userHandler,
	}
}

func (server *Server) Start(conf config.Config) {
	mux := http.NewServeMux()

	middleware := middlewares.NewManager()
	middleware.Use(middlewares.Logger)

	server.userHandler.RegisterRoutes(mux, middleware)

	h := middleware.WrapMux(mux)

	port := ":" + conf.HttpPort

	fmt.Println("http://localhost:", conf.HttpPort)
	http.ListenAndServe(port, h)
}
