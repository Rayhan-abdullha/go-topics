package product

import (
	"server/rest/middleware"
)

type Handler struct {
	middlewares *middleware.MiddlewareType
	svc         Service
}

func NewProductHandler(
	middlewares *middleware.MiddlewareType,
	svc Service) *Handler {
	return &Handler{
		middlewares: middlewares,
		svc:         svc,
	}
}
