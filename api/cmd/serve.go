package cmd

import (
	"article/config"
	"article/rest"
	"article/rest/handlers/user"
)

func Serve() {
	conf := config.GetConfig()
	userHandler := user.NewHandler()
	server := rest.NewServer(userHandler)
	server.Start(*conf)
}
