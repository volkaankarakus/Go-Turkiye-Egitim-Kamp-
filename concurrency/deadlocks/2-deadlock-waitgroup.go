package main

import (

)

// ** diyelim ki yanlıs sayida bekleyecegimiz go routinei verdik diyelim
// func main() {
// 	// ** waitGroup for synchronization
// 	waitGroup := sync.WaitGroup{}

// 	waitGroup.Add(2)

// 	go func() {
// 		fmt.Println("hello from go-routine")
// 		waitGroup.Done()
// 	}()

// 	waitGroup.Wait()
// 	fmt.Println("hello concurrency")
// }

// ** bu da bir DEADLOCK sorunudur.
// **  biz daha fazla kisiyi bekliyoruz, ama bunu DONE'a cekicek kimse yok.
