package singleton

import (
	"fmt"
	"sync"
)

type User struct {
	Name  string
	Email string
}

var (
	instance *User
	once     sync.Once
)

func GetInstance() *User {
	once.Do(func() {
		instance = &User{
			Name:  "Rayhan",
			Email: "rayhan@gmail",
		}
	})
	return instance
}

func SingleTon() {
	u1 := GetInstance()
	u2 := GetInstance()

	u2.Name = "Abir"

	fmt.Println(u1)
	fmt.Println(u2)
}
