package practice

import "fmt"

func arrExample() {
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println(arr)
}
func arrExample2() {
	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	// arr1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	arr[4] = 200
	// arr1[3] = 410
	fmt.Println(arr[:2]) //
	fmt.Println(arr[2:]) // skip first 2 value
}
func arrIterate() {
	arr := [5]int{1, 2, 8, 3, 4}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }
	for i, val := range arr {
		fmt.Println(i, " = ", val)
	}
}
func productArr() {
	type Product struct {
		ID   int
		Name string
	}
	shoppingList := [4]Product{
		{ID: 1, Name: "Camera"},
		{ID: 2, Name: "Laptop"},
		{ID: 3, Name: "Phone"},
	}
	fmt.Println(shoppingList)

	for i, val := range shoppingList {
		pd := val
		if pd.Name == "" {
			shoppingList[i] = Product{ID: 4, Name: "Microphone"}
		}
	}
	fmt.Println(shoppingList)

}

func Array() {
	// arrExample()
	// arrExample2()
	// arrIterate()
	productArr()
}
