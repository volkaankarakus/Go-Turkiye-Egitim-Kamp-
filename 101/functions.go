package main

import (
	// "fmt"
	// "time"
)

// func main() {
// 	timer(time.After(5 * time.Second))
// }

// // ** DEFER
// // ** defer; func bittikten sonra calismaya baslar ve kod satiri olarak tersten calisir.

// func timer(c <-chan time.Time, message ...string) {
// 	defer fmt.Println("bir sonraki asamaya geciliyor..")
// 	defer fmt.Println("timer bitti")

// 	for {
// 		select {
// 		case <-c:
// 			return
// 		default:
// 			fmt.Println(time.Now(), message)
// 			time.Sleep(1 * time.Second)
// 		}   
// 	}
// }
