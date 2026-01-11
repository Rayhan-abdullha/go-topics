package product

import (
	"server/domain"
	userHandler "server/rest/handler/product"
)

type Service interface {
	userHandler.Service // embed existing service interface
}

type ProductRepo interface {
	Create(*domain.Product) (*domain.Product, error)
	List() ([]*domain.Product, error)
}
