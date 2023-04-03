// ** 6'daki problemi waitGroupsuz cözmeye calisalim

package main

// func main() {
// 	bufferedChannel := make(chan int, 5)
// 	done := make(chan bool)

// 	go func(bufChan chan int, done chan bool) {
// 		for val := range bufChan {
// 			fmt.Println(val)
// 		}
// 		fmt.Println("channel closed")
// 		done <- true
// 	}(bufferedChannel, done)

// 	bufferedChannel <- 1
// 	bufferedChannel <- 2
// 	bufferedChannel <- 3
// 	bufferedChannel <- 4
// 	bufferedChannel <- 5
// 	bufferedChannel <- 6
// 	bufferedChannel <- 7
// 	bufferedChannel <- 8
// 	bufferedChannel <- 9
// 	close(bufferedChannel) // ** bufferChannel otomatik kendisi dispose olur ama biz yine de kapatalim

// 	// ** wait go-routine to finish
// 	<-done // mainde de done channelindan data okumaya calisiyoruz.

// 	fmt.Println("main closed")
// 	// OUTPUT : 1 2 3 4 5 6 7 8 9 channel closed main closed
// }

// ** unbuffered channellarda yazma ve okuma isleri blockingdir demistik.
// ** 31. satirda done'a data basana kadar main go-routine blocklaniyor.
