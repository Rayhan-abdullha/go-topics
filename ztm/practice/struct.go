package practice

import "fmt"

type Data struct {
	field string
	a     int
	b     int
}

func Struct() {
	data := Data{
		field: "coder",
		a:     3,
		b:     4,
	}
	word := data.field
	data.a = 1000
	a, b := data.a, data.b
	fmt.Println(word)
	fmt.Println(a, b)
}
