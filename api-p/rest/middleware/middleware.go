package middleware

import "server/config"

type MiddlewareType struct {
	cnf *config.Config
}

func NewMiddleware(cnf *config.Config) *MiddlewareType {
	return &MiddlewareType{
		cnf: cnf,
	}
}
