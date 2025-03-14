package basic

import "fmt"

func DataStruct() {
	str := "Hello, Go!"
	fmt.Println("str[0]: " + fmt.Sprint(str[0]) + " ('H' in ASCII)")

	// Array (fixed size)
	arr := [3]int{1, 2, 3}
	fmt.Println("arr: " + fmt.Sprint(arr))

	// Slice (dynamic size)
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Println("slice: " + fmt.Sprint(slice))

	// Make & Map
	m := make(map[string]int)
	m["age"] = 30
	fmt.Println("m[\"age\"]: " + fmt.Sprint(m["age"]))
}
