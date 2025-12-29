package main

import (
	"fmt"
)

type ServerRequest map[string]int

func addServer(ser ServerRequest, serverName string) {
	ser[serverName] = 0
}
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
				next = "server4"
			case "server4":
				next = "server1"
			}
		}
	}
	fmt.Println(ser)
}

func main() {
	server := ServerRequest{}
	addServer(server, "server1")
	addServer(server, "server2")
	addServer(server, "server3")
	addServer(server, "server4")

	roundRobinAlgo(server, 100)
}
