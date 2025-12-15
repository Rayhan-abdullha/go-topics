package user

import "article/rest/middlewares"

type Handler struct {
	middleware *middlewares.Middleware
}

func NewHandler(middleware *middlewares.Middleware) *Handler {
	return &Handler{
		middleware: middleware,
	}
}
