package main

import (
	"fmt"
	"pattern/singleton"
)

func main() {
	user := singleton.User{
		Name:  "Rayhan",
		Email: "rayhan@gmail",
	}
	singleton.GetInstance(user)
	fmt.Println(user.GetUser())
}
