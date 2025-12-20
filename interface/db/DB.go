package db

type User struct {
	ID   int
	Name string
}
type UserRepo interface {
	Find()
	Create()
}
type userRepo struct {
	userList []User
}

func NewUserRepo() UserRepo {
	return &userRepo{
		userList: []User{},
	}
}
func (r *userRepo) Find() {
	for _, user := range r.userList {
		println("User ID:", user.ID, "Name:", user.Name)
	}
}
func (r *userRepo) Create() {
	r.userList = append(r.userList, User{ID: 1, Name: "Alice"})
}

func DB() {
	repo := NewUserRepo()
	repo.Create()
	repo.Create()
	repo.Find()
}
