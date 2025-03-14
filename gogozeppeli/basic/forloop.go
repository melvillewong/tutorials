package basic

import "fmt"

func Forloop() {
	i := 0

	// while-ish loop
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// Modern ranged for loop
	for j := range [5]int{} {
		fmt.Println(j)
	}

	// Infinite loop
	for {
		fmt.Println("Infinite")
		break
	}

	// loop for data struct
	nums := []int{1, 2, 3, 4, 5}
	for i, num := range nums {
		fmt.Println("nums[" + fmt.Sprint(i) + "] = " + fmt.Sprint(num))
	}

	users := map[string]string{"name": "Alice", "age": "30", "city": "Hong Kong"}
	for key, val := range users {
		fmt.Println(key + ": " + val)
	}
}
