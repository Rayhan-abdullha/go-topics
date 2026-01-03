package main

import (
	"fmt"
)

type User struct {
	name string
}

// func (u *User) getUser() {
// 	fmt.Println(u)
// 	u.name = "developer"
// }

func main() {
	// user := User{name: "coder"}
	// user.getUser()
	// fmt.Println(user)
	// a := [5]int{1, 22, 3, 4, 3}
	// b := []int{1, 22, 3, 4, 3}
	// b[1] = 100
	// a[1] = 200
	// fmt.Println("a", a)
	// fmt.Println("b", b)
	// a := &user
	// a.name = "changed"
	// fmt.Println(user)

	u := make([]int, 3, 5)
	fmt.Println(u)

}
