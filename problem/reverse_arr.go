package problem

import "fmt"

// arr = [1, 2, 3, 4, 5]
func Reverse_arr(arr []int) {
	j := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
	}
	fmt.Println(arr)
}
