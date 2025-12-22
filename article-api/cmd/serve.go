package cmd

import (
	"article/config"
	"article/infra/db"
	"article/rest"
	"article/rest/handlers/user"
	"article/rest/middlewares"
	"fmt"
	"os"
)

func Serve() {
	conf := config.GetConfig()
	dbCon, err := db.NewConnection()
	if err != nil {
		fmt.Println("connection error", err)
		os.Exit(1)
	}
	mid := middlewares.NewMiddleware(&conf)
	server := rest.NewServer(
		&conf,
		user.NewHandler(mid, dbCon),
	)
	server.Start(conf)
}
