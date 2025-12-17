package main

import "fmt"

// import "golang/interface/auth"
type User interface{}

func Auth(v User) {
	fmt.Println(v)
}
func main() {
	// auth.Auth()

	Auth(3)
}
