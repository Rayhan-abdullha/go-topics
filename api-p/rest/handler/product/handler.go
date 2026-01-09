package product

import (
	"server/repo"
	"server/rest/middleware"
)

type Product struct {
	middlewares *middleware.MiddlewareType
	productRepo repo.ProductRepo
}

func NewProductHandler(
	middlewares *middleware.MiddlewareType,
	productRepo repo.ProductRepo) *Product {
	return &Product{
		middlewares: middlewares,
		productRepo: productRepo,
	}
}
