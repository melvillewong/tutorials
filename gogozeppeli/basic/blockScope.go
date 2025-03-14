package basic

import "fmt"

// Only way to declare var globally
var global int = 10

func Scope() {
	x := 10 // 變數 x 的作用範圍是 main 函式的區塊
	if x > 5 {
		y := 20               // 變數 y 的作用範圍是 if 區塊
		fmt.Println("x =", x) // 可以訪問 x，因為 x 是在外部區塊宣告的
		fmt.Println("y =", y) // 可以訪問 y，因為 y 是在該區塊內宣告的
	}
	// fmt.Println("y =", y) // 編譯錯誤：y 在此無法訪問
}
