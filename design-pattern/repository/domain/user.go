// domain/user_repository.go
package domain

type UserRepository interface {
	FindByID(id int) (*User, error)
	Save(user *User) error
}
