package main

import (
	"fmt"
	"sync"
)

// "server/cmd"

func main() {
	// cmd.Serve()

	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 1
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		data := <-ch
		fmt.Println(data)
	}()
	wg.Wait()
	fmt.Println("done")
}
