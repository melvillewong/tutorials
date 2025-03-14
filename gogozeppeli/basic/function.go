package basic

import "fmt"

func concat(a string, b string) string {
	return a + b
}

// mutliple type
func add(a, b int, c string) int {
	fmt.Println(c)
	return a + b
}

// return mutliple values
func getName() (string, string) {
	return "Melville", "Wong"
}

// Naming return value is easy to understand what're returning
func getCoords() (x, y int) {
	return // automatically return x and y with uninitialised value
}

// better to explicitly return the values even it works if u dont

func Function() {
	// ignore a return value
	firstName, _ := getName()
	fmt.Println(firstName)
}
