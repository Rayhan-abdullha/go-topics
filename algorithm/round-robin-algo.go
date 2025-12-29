package main

import (
	"fmt"
)

type ServerRequest map[string]int

func roundRobinAlgo(ser ServerRequest, req int) {
	var next string = "server1"
	for i := 1; i <= req; i++ {
		if len(ser) == 1 {
			ser[next]++
		} else {
			ser[next]++
			switch next {
			case "server1":
				next = "server2"
			case "server2":
				next = "server3"
			case "server3":
				next = "server1"
			}
		}
	}
	fmt.Println(ser)
}

func main() {
	server := ServerRequest{
		"server1": 0,
		"server2": 0,
		"server3": 0,
	}
	roundRobinAlgo(server, 100)
}
