package main

import (
	"fmt"
	"my-go/mathlib"
)

type Nums struct {
	x int
}

func main() {
	fmt.Println("main package")
	mathlib.Add(3, 3)
}
