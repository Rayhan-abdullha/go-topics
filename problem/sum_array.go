package problem

import "fmt"

func SumArray(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	fmt.Println(sum)
	return sum
}
func MillionLoop(c int) {
	sum := 0
	for i := 0; i < c; i++ {
		sum += i
	}
}
