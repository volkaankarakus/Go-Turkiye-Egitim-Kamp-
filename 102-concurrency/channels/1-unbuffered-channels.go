package main


// ** Go'nun felsefesi : memoryi etrafa dagıtarak haberlesmemek
// ** Do not communicatate by sharing memory instead share memory with communication

// ** POINTERLARI ETRAFA SACMA, CHANNEL OLUSTUR VE BUNLARI GO-ROUTINELERE VER
// ** channellar dogası geregi blocking islemler yapabilen, okuma ve yazmanın blocking oldugu typetır.

// ** channellar MAKE ile olusturulur.
// ** channel ana kodlarını incelersek channelin lock'lanan bir array oldugunu goruruz.

// ** Channellar ikiye ayrılır : Buffered ve Unbuffered
// ** UNBUFFERED : Her okuma ve yazma islemi blockingdir. Es zamanlı olarak sadece 1 data yazıp, 1 data okuyabiliriz.
// ** BUFFERED : Buffered'a size veriyoruz, verdigimiz size'a kadar blocking olmuyo. Bu sizedan sonra blocking oluyor.

// func main() {
// 	// ** channel initialization
// 	// ** CHANNELLAR MAKE ILE TANIMLANIR VE ARRAY GIBI ICERISINDE TURUYLE TANIMLANIR
// 	unbufferedChannel := make(chan int)

// 	// ** channel decleration
// 	var unbufferedChannel2 chan int // bu printlenince nil doner

// 	fmt.Println(unbufferedChannel)  // 0x140000220c0
// 	fmt.Println(unbufferedChannel2) // <nil>

// 	// ** only can read from channel
// 	go func(unbufferedChan <-chan int) {
// 		// data gelene kadar go-routine bloklanir.(default olarak)
// 		value := <-unbufferedChan
// 		fmt.Println(value) // Output : 1
// 	}(unbufferedChannel)

// 	unbufferedChannel <- 1
// }
