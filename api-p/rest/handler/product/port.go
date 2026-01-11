package product

import "server/domain"

type Service interface {
	Create(*domain.Product) (*domain.Product, error)
	List() ([]*domain.Product, error)
}
