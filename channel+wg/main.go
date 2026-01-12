package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name string
	Age  int
}

func main() {
	ch := make(chan User)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("Sending data from goroutine 1")
		defer wg.Done()
		ch <- User{Name: "Alice", Age: 30}
	}()

	wg.Add(1)
	go func() {
		fmt.Println("Receiving Data from goroutine ")
		defer wg.Done()
		data := <-ch
		fmt.Println(data)
	}()
	wg.Wait()
}
