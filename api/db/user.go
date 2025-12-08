package db

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var DB []User

func CreateUser(user User) {
	DB = append(DB, user)
}
func GetUser() []User {
	return DB
}
