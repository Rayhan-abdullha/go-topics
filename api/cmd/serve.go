package cmd

import (
	"article/config"
	"article/rest"
	// "article/router"
)

func Serve() {
	conf := config.GetConfig()
	rest.Start(conf)
}
