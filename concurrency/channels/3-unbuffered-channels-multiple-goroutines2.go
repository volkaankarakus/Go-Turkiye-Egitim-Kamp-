package main

import (
	"fmt"
	"time"
)

func main() {
	unbufferedChannel := make(chan int)

	// ** reader go-routine
	go func(unbufChan chan int) {
		value := <-unbufChan
		fmt.Println(value)
	}(unbufferedChannel)

	// ** writer go-routine
	go func(unbufChan chan int) {
		unbufChan <- 1
	}(unbufferedChannel)

	time.Sleep(1 * time.Second) // ******
	fmt.Println("hello channels")

}

// ** 2-unbuffereddaki UNDETERMINISTIC'ligi sleep atarak handle edebiliriz.
