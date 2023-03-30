package main

import "fmt"

// ** 1. YOL : var ile tanimlama
var array1 [3]int                  // [0 0 0]
var array2 = [5]int{1, 2, 3, 4, 5} // [1 2 3 4 5]

// func main() {
   

// }

// ** 2. YOL : func icinde make ile tanimlama
func funcIcindeArrayTanimi() {
	array3 := make([]int, 3)
	array3[1] = 3
	fmt.Println(array3) // [0 3 0]
}
