package user

import (
	"article/rest/middlewares"

	"github.com/jmoiron/sqlx"
)

type Handler struct {
	middleware *middlewares.Middleware
	dbConf *sqlx.DB
}

func NewHandler(middleware *middlewares.Middleware, dbConf *sqlx.DB) *Handler {
	return &Handler{
		middleware: middleware,
		dbConf: dbConf,
	}
}
