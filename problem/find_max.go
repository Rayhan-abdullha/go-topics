package problem

import "fmt"

func FindMax(arr []int) int {
	max := arr[1]
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		if max < temp {
			max = temp
		}
	}
	fmt.Println(max)
	return max
}
