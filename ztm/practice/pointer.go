package practice

import "fmt"

func incrementVal(val *int){
	*val += 1
}
func pointerExample() {
	val := 10
	// valuePtr := &val
	// fmt.Println(*valuePtr)
	incrementVal(&val) // 11
	incrementVal(&val) // 12
	incrementVal(&val) // 13
	fmt.Println(val)
}

func Pointer() {
	pointerExample()
}
