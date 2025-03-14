package basic

import "fmt"

func Format() {
	name := "Alice"
	age := 21

	// print out '%' sign
	fmt.Printf("%%s = %s\n", name)

	// '%t' = boolean
	fmt.Println("%t: boolean")
	// '%T' = variable type
	fmt.Println("%T: variable type")
	// '%v' = common type (autoly takes the type of the var itself)(lazy method)
	fmt.Println("%v: no specific type (use type of the var itself)")

	// format convention
	fmt.Println("Convent int to string: " + fmt.Sprintf("%d", age))
}
