package day1

import "fmt"

type User struct {
	Name string
	Age  int
}

func update(u *User) {
	u.Age = 21
}
func Main() {
	usr := &User{
		Name: "Abir",
		Age:  20,
	}
	a := usr
	update(usr)
	fmt.Println(usr)
	a.Age = 22
	fmt.Println(usr)
}
