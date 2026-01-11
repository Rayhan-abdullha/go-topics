package product

import "server/domain"

type service struct {
	productRepo ProductRepo
}

func NewService(pr ProductRepo) Service {
	return &service{
		productRepo: pr,
	}
}

func (svc *service) Create(p *domain.Product) (*domain.Product, error) {
	pd, err := svc.productRepo.Create(p)
	if err != nil {
		return nil, err
	}
	if pd == nil {
		return nil, nil
	}
	return pd, nil
}
func (svc *service) List() ([]*domain.Product, error) {
	pds, err := svc.productRepo.List()
	if err != nil {
		return nil, err
	}
	return pds, nil
}
