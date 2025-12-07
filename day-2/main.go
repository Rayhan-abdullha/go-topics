package day2

import "fmt"

// interface
type Animal interface {
	Speak()
}

// struct
type Dog struct{}
type Cat struct{}

// implement method
func (Dog) Speak() {
	fmt.Println("Bark Bark...")
}
func (Cat) Speak() {
	fmt.Println("Meow Meow...")
}
func great(f Animal) {
	f.Speak()
}
func Main() {
	d := Dog{}

	// c := Cat{}
	// c.Speak()

	great(d)
	great(Cat{})
}
