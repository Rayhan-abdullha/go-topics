package practice

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarder      bool
}
type Bus struct {
	FrontSeat Passenger
}

func Struct1() {
	casey := Passenger{"Casey", 1, false}
	fmt.Println(casey)
	var (
		bill = Passenger{"Bill", 2, false}
		jeff = Passenger{"Jeff", 3, false}
	)
	fmt.Println(bill, jeff)
	bus := Bus{
		jeff,
	}
	fmt.Println(bus)
}
