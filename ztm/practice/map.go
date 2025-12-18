package practice

import "fmt"

type Obj struct {
	Name string
	Age  int
}

func mapExample() {
	// obj := map[string]int{
	// 	"one":   1,
	// 	"two":   2,
	// 	"three": 3,
	// }
	// obj["four"] = 4
	// fmt.Println(obj)

	user := make(map[string]string)
	user["name"] = "Abir"
	user["email"] = "VY2kP@example.com"
	user["password"] = "1234"
	user["password1"] = "12345444"
	user["password2"] = "1234777"
	delete(user, "password")
	if email, ok := user["email"]; ok {
		fmt.Println(email)
	}

	// iterate obj value
	for _, val := range user {
		fmt.Println(val)
	}

	
}

func Map() {
	mapExample()
}
