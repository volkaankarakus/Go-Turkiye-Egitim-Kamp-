package main


// func main(){
// 	channel1 := make(chan int,1)
// 	channel1 <- 1

// 	channel2 := make(chan int,2)
// 	channel2 <- 2

// 	select{
// 	case channel1Value := <- channel1:
// 		fmt.Println(channel1Value)
// 	case channel2Value := <- channel2 :
// 		fmt.Println(channel2Value)
// 	}
// }

// ** printte ya 1 ya da 2'yi basıyor. iki printi de aynı anda basmıyor.
// ** select icerisinde multiple channel calistirdigimizda rastgele bir channel secer.
// ** Bu sorunu asabilmek icin select'i for dongusune sokarız ki bir sonraki channel'i case'e dusursun.

// func main() {
// 	channel1 := make(chan int, 1)
// 	channel1 <- 1

// 	channel2 := make(chan int, 2)
// 	channel2 <- 2

// 	for {
// 		select {
// 		case channel1Value := <-channel1:
// 			fmt.Println(channel1Value)
// 		case channel2Value := <-channel2:
// 			fmt.Println(channel2Value)
// 		}
// 	}
// }

// OUTPUT : 2 1 fatal error: all goroutines are asleep - deadlock!
// ** bunun nedeni 2 channeldan birisini seçti ve bastı. sonra selectten dısarı cıkıp diger channeli secti ve bastı.
// ** sonra tekrar geldiginde herhangi bir channela data basan kimse yok. o yüzden deadlock olusur.

// ** Bu sorunu FLAG tanımlayarak cozebiliriz
// func main() {
// 	channel1 := make(chan int, 1)
// 	channel1 <- 1

// 	channel2 := make(chan int, 2)
// 	channel2 <- 2

// 	var done bool
// 	for !done {
// 		select {
// 		case channel1Value := <-channel1:
// 			fmt.Println(channel1Value)
// 		case channel2Value := <-channel2:
// 			fmt.Println(channel2Value)
// 		default:
// 			done = true
// 		}
// 	}
// } // OUTPUT : 2 1 ya da 1 2 .. sorun cözüldü.
