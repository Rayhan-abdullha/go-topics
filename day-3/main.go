package day3

import "fmt"

type List struct {
	Arr [6]any
}

func (l *List) append(val interface{}) {
	if l.Arr[0] == nil {
		l.Arr[0] = val
	} else {
		for i, val := range l.Arr {
			if val == nil {
				l.Arr[i] = val
				return
			}
		}
	}
}
func (a *List) getList() {
	fmt.Println(a.Arr)
}
func Main() {
	arr := List{}
	arr.append(1)
	arr.append(2)
	arr.getList()
}
