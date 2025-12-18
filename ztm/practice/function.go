package practice

import "fmt"

var (
	a      = 10
	b int  = 41
	c bool = true
)

func getData() (data string, err error) {
	if data == "" {
		err = fmt.Errorf("data is empty")
	} else {
		err = nil
	}
	return data, err
}
func Function() {
	data, err := getData()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}
}
