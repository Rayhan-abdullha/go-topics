package main

import (
	"fmt"
)

func increment() func() {
	n := 0
	count := func() {
		n++
	}
	return count
}

func main() {
	count := increment()
	count()
	count()
	fmt.Println(count)
}
