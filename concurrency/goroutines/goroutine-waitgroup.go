package main

import (

)

// ** Go,Go-routine ve concurrency dedigimizde aklimiza gelmesi gereken temel seylerden birisi : WAITGROUP
// ** Waitgroup, belirli sayida goroutine'i beklememiz icin olusturulmus bir tiptir.

// func main() {
// 	waitGroup := sync.WaitGroup{}
// 	waitGroup.Add(1) // ** 1 tane go routine beklemek istiyorum diyorum.

// 	go func() {
// 		fmt.Println("hello from goroutine")
// 		waitGroup.Done()
// 	}()

// 	waitGroup.Wait()

// 	fmt.Println("hello from main")
// }
