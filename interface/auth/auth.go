package auth

import "fmt"

type Authenticator interface {
	Login(username, password string) bool
}
type UserService struct{}

func (u UserService) Login(username, password string) bool {
	return username == "admin" && password == "1234"
}

type GoogleAuth struct{}

func (g GoogleAuth) Login(username, password string) bool {
	if username == "" || password == "" {
		return false
	}
	return username == "admin2" && password == "1234"

}

func Authenticate(a Authenticator, username, password string) {
	if a.Login(username, password) {
		fmt.Println("Login success")
	} else {
		fmt.Println("Credentials are wrong")
	}
}

func Auth() {
	Authenticate(UserService{}, "admin", "1234")
	Authenticate(GoogleAuth{}, "admin2", "1234")

}
