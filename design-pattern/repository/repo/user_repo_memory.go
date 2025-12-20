package repo

import (
	"pattern/repository/domain"
	"sync"
)

type InMemoryUserRepository struct {
	mu     sync.RWMutex
	data   map[int]*domain.User
	nextID int
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		data:   make(map[int]*domain.User),
		nextID: 1,
	}
}

func (r *InMemoryUserRepository) Save(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextID
	r.nextID++

	r.data[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) GetAll() ([]*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]*domain.User, 0, len(r.data))
	for _, user := range r.data {
		users = append(users, user)
	}
	return users, nil
}
