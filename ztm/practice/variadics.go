package practice

import "fmt"

func sum(num ...int) {
	sum := 0
	for _, val := range num {
		sum += val
	}
	fmt.Println(sum)
}
func Variadics() {
	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7}
	all := append(a, b...)
	sum(all...)
}