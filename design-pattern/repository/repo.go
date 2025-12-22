package repository

import (
	"fmt"
	repository "pattern/repository/repo"
	service "pattern/repository/services"
)

func RepoPattern() {
	userRepo := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepo)

	userService.CreateUser("Bob")

	users, _ := userService.ListUsers()
	fmt.Println("All users:")
	for _, u := range users {
		fmt.Println(u)
	}
}
