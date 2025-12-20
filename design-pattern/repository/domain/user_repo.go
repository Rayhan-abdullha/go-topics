package domain

type UserRepository interface {
	Save(user *User) error
	GetAll() ([]*User, error)
}
