package practice

import (
	"errors"
	"fmt"
)

type User struct {
	ID int
	Name  string
	Email string
}
type DB struct {
	Users []User
}
func (db *DB) FindLast() (usr *User, err error) {
	if len(db.Users) == 0 {
		return nil, errors.New("User is empty!")
	}
	return &db.Users[len(db.Users) -1], nil
}
func (u *DB) findUser() (user []User, err error) {
	if len(u.Users) == 0 {
		return nil, errors.New("users is empty")
	}
	return u.Users, nil
}
func (u *DB) createUser(usr User) {
	last, err := u.FindLast()
	if err != nil {
		usr.ID = 1
		u.Users = append(u.Users, usr)
	} else {
		usr.ID = last.ID+1
		u.Users = append(u.Users, usr)
	}
}

func ReceiverFunction() {
	db := DB{}
	user := User{Name: "rayhan", Email: "rayhan@gmail.com"}
	db.findUser()
	db.createUser(user)
	// db.createUser(user)
	// db.createUser(user)
	// usr := db.findUser()
	fmt.Println(user)
}