package main

import (
)

// func main() {
// 	// ** channel initialization
// 	bufferedChannel := make(chan int, 5)

// 	go func(bufChan chan int) {
// 		for {
// 			val := <-bufChan
// 			fmt.Println(val)
// 		}
// 	}(bufferedChannel)

// 	bufferedChannel <- 1
// 	bufferedChannel <- 2
// 	bufferedChannel <- 3
// 	bufferedChannel <- 4
// 	bufferedChannel <- 5
// 	bufferedChannel <- 6
// 	bufferedChannel <- 7
// 	bufferedChannel <- 8
// 	bufferedChannel <- 9

// 	// OUTPUT : 1 2 3 4 5 6 7
// 	// bir sonraki OUTPUT : 1 2 3 4 5 6 7 8 9
// 	// ** UNDETERMINISTIC OUTPUT
// 	// ** Butun Outputu bu go-routine'in yazicagi garanti degil.Cunku bu go-routinein bitmesini beklemiyoruz.
// 	// ** 26. satirdan sonra main go-routine bittigi icin arka plandaki butun go-routineler de ölüyor.(yani islemini bitirmeyebilirler)
// 	//   ** beklemek istiyorsak waitGroup kullanabiliriz.
// }
