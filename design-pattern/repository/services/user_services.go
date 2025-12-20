package service

import "pattern/repository/domain"

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name string) (*domain.User, error) {
	user := &domain.User{Name: name}
	err := s.repo.Save(user)
	return user, err
}

func (s *UserService) ListUsers() ([]*domain.User, error) {
	return s.repo.GetAll()
}
