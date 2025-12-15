package middlewares

import "article/config"

type Middleware struct {
	Cnf *config.Config
}

func NewMiddleware(cnf *config.Config) *Middleware {
	return &Middleware{
		Cnf: cnf,
	}
}
