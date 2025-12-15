package cmd

import (
	"article/config"
	"article/rest"
	"article/rest/handlers/user"
	"article/rest/middlewares"
)

func Serve() {
	conf := config.GetConfig()
	mid := middlewares.NewMiddleware(&conf)
	server := rest.NewServer(
		&conf,
		user.NewHandler(mid),
	)
	server.Start(conf)
}
