package main

import (

)

// ** FAN-IN : Birden fazla channela input geliyordur. Biz bu channellari biraraya getirip tek bir channel olarak dısarıya vermek istiyoruzdur.
// ** 3 tane channelimiz ve 1 tane de output channelimiz olsun. Islemlerimizin bittigini belirten bir tane de "done" channelimiz olsun.

// func main() {
// 	channel1 := make(chan int)
// 	channel2 := make(chan int)
// 	channel3 := make(chan int)

// 	outputChannel := make(chan int, 1000)

// 	done := make(chan int)

// 	go func() {
// 		for {
// 			fmt.Println("worket waits...")
// 			value, open := <-outputChannel
// 			if !open {
// 				break
// 			}
// 			fmt.Println(value)
// 		}
// 		done <- 1
// 	}()

// 	go func() {
// 		for {
// 			channel1 <- 1
// 			time.Sleep(1 * time.Second)
// 		}
// 	}()

// 	go func() {
// 		for {
// 			channel2 <- 2
// 			time.Sleep(1 * time.Second)
// 		}
// 	}()

// 	go func() {
// 		channel3 <- 3
// 		time.Sleep(1 * time.Second)
// 	}()

// 	fanIn([]chan int{channel1, channel2, channel3}, outputChannel)
// 	<-done
// }

func fanIn(inChans []chan int, outputChan chan int) {
	for _, channel := range inChans {
		go func(c chan int) {
			for {
				val, open := <-c
				if !open {
					break
				}
				outputChan <- val
			}
		}(channel)
	}
}

// ** bilgi olarak bil. Algoritmayı anlamaya calis
// ** worker waits... 2 worker waits... 1 worker waits... 3 worker waits... 1 worker waits... 2
