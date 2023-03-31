package main


// ** SELECT-CASE CHANNELS
// ** select-case ile bir channel'a veri girisini bekletebiliyoruz

// func main() {
// 	channel1 := make(chan int, 1)
// 	channel1 <- 1

// 	select {
// 	case channel1Value := <-channel1:
// 		fmt.Println(channel1Value)
// 	}

// }

// ** case'de diyoruz ki : bu channeli dinle, eger bu channela veri gelirse bunu outputa yaz.
// ** select case ile bu sekilde channel dinlenebiliyor.
