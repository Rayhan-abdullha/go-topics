package rest

import (
	"article/config"
	middleware "article/rest/middlewares"
	"article/rest/routes"
	"fmt"
	"net/http"
)

func Start(conf config.Config) {
	// global middleware
	m := middleware.NewManager()
	m.Use(
		middleware.CorsAndPrefix,
		middleware.Logger,
		middleware.AuthMiddleware,
	)

	mux := http.NewServeMux()

	h := m.WrapMux(mux)

	routes.InitialRoute(mux, m)
	port := ":" + conf.HttpPort
	fmt.Println("Listening on port ", port)
	http.ListenAndServe(port, h)
}
