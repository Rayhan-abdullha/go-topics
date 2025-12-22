package main

import "fmt"

type DB struct {
	users []User
	products []Product
}
func NewCreateDB() *DB{
	return &DB{
		users: []User{},
		products: []Product{},
	}
}

type Db_Operation interface {
	Create()
	Find()
} 

// user
type User struct {
	ID int
	Name string
}
func (u User) Create() {
	fmt.Println("created user")
}
func (u User) Find(){
	fmt.Println([]User{
		{ID: 1, Name: "abir"},
		{ID: 1, Name: "rayhan"},
	})
}

// product
type Product struct {
	ID int
	Name string
	Category string
}
func (u Product) Create() {
	fmt.Println("created product")
}
func (u Product) Find(){
	fmt.Println([]Product{
		{ID: 1, Name: "camera", Category: "b"},
		{ID: 1, Name: "phone", Category: "a"},
	})
}

func operation(i Db_Operation) {
	i.Create()
}



func main() {
	// create database
	// db := NewCreateDB()

	user := User{ID: 3, Name: "messi"}
	operation(user)
	
	pd := Product{ID: 33, Name: "laptop", Category: "c"}
	operation(pd)
	
}
