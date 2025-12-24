package problem

import "fmt"

func sum(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] += 10
	}
}

func Problem() {
	// SumArray([]int{1, 2, 3, 4, 5})

	// FindMax([]int{1, 2, 33, 21, 34, 22})

	arr := []int{1, 2, 3, 4, 5}
	// fmt.Println("before", arr)
	// Reverse_arr(arr)
	// fmt.Println("after: ", arr)

	sum(arr)
	fmt.Println(arr)
}
