package food

import "fmt"

type Prepare interface {
	PrepareDishes()
}
type Chicken string
type Salad string

func (c Chicken) PrepareDishes() {
	fmt.Println("Chicken is Prepared")
}
func (s Salad) PrepareDishes() {
	fmt.Println("Salad is Prepared")
}

func prepareDishes(dishes []Prepare) {
	fmt.Println("Preparing Dishes...")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		dish.PrepareDishes()
	}
}
func Dish() {
	chicken := Chicken("Chicken")
	salad := Salad("Salad")
	arr := []Prepare{chicken, salad}
	
	prepareDishes(arr)

}
