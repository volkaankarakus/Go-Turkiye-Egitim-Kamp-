package main

import (

)

// func main() {
// 	//** channel initialize
// 	unbufferedChannel := make(chan int)

// 	// ** unbufferedChannel'dan data okuyan go-routine
// 	go func(unbufChan chan int) {
// 		value := <-unbufChan
// 		fmt.Println(value)
// 	}(unbufferedChannel)

// 	// ** unbufferedChannela data yazan go-routine
// 	go func(unbufChan chan int) {
// 		unbufChan <- 1
// 	}(unbufferedChannel)

// 	fmt.Println("hello multiple go-routines")
// }

// ** Runladıgımızda ekrana datayı basmadı sadece "hello go-routines" bastı.
// ** Schedulerdan dolayı bunun outputu yine UNDETERMINISTIC'tir.
