package singleton

type User struct {
	Name  string
	Email string
}

var instance *User

func GetInstance(user User) *User {
	if instance == nil {
		instance = &user
	}
	return instance
}
func (u User) GetUser() User {
	return u
}
