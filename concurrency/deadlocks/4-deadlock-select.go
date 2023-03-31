package main

import "fmt"

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go func(c chan int) {
		c <- 1
	}(channel1)

	go func(c chan int) {
		c <- 2
	}(channel2)

	for {
		select {
		case channel1Value := <-channel1:
			fmt.Println(channel1Value)
		case channel2Value := <-channel2:
			fmt.Println(channel2Value)
		}
	}
}

// ** daha once bundan bahsetmistik. for sürekli oldugu icin deadlock yer.
