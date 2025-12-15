package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
}

func NewUser(name string) *User {
	return &User{
		name: name,
	}
}
func (User) log(t time.Time) {
	fmt.Println(time.Since(t))
}
func (u *User) myFunc() {
	t := time.Now()
	fmt.Println(u.name)
	// write wite code
	time.Sleep(time.Second * 2)
	u.log(t)
}
func main() {
	a := NewUser("abir")
	b := NewUser("rayhan")
	a.myFunc()
	b.myFunc()
}
