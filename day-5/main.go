package day5

import "fmt"

type User struct {
	Name  string
	Email string
}
type Users struct {
	Users []User
}

func (U *Users) CreateUser(user User) {
	U.Users = append(U.Users, user)
}

func Main() {
	user := User{
		Name:  "Rayhan",
		Email: "rayhan@gmail",
	}
	users := Users{}
	users.CreateUser(user)
	users.CreateUser(user)
	users.CreateUser(user)
	fmt.Println(users.Users)
}
