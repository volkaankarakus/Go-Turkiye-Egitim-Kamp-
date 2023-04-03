package main

// ** BUFFERED : Buffered'a size veriyoruz, verdigimiz size'a kadar blocking olmuyo. Bu sizedan sonra blocking oluyordu.

// func main() {
// 	// ** channel initialization
// 	bufferedChannel := make(chan int, 5)

// 	bufferedChannel <- 1
// 	bufferedChannel <- 2
// 	bufferedChannel <- 3
// 	bufferedChannel <- 4
// 	bufferedChannel <- 5

// 	fmt.Println(<-bufferedChannel) // 1
// 	fmt.Println(<-bufferedChannel) // 2
// 	fmt.Println(<-bufferedChannel) // 3
// 	fmt.Println(<-bufferedChannel) // 4
// 	fmt.Println(<-bufferedChannel) // 5
// }
