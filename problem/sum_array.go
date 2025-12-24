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
