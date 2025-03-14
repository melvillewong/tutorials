package basic

import "fmt"

func Switch() {
	fruit := "apple"

	// Default: break after executing every case

	// fallthrough executes 1 next consecutive case
	switch fruit {
	case "apple":
		fmt.Println("This is an apple")
		fallthrough
	case "banana":
		fmt.Println("This is a banana")
	case "orange":
		fmt.Println("This is an orange")
	}
}
