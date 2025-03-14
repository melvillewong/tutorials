package basic

import "fmt"

func Variable() {
	// quick init
	name := "Melville"
	fmt.Println(name)

	// init
	var name1, name2 string
	name1 = "Charlotte"
	name2 = "Anson"
	fmt.Println(name1, name2)

	// multi quick init
	name3, name4 := "Bamboo", "Chapman"
	fmt.Println(name3, name4)

	// multi init
	var (
		name5 string
		age   int
	)
	name5, age = "Matfrey", 21
	fmt.Println(name5, age)
}
