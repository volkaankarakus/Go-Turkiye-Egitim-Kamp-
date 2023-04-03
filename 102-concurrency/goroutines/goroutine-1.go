package main


// ** asagidaki fonksiyonun basında go yazmasaydi kod satirlarinin sirasina göre kod calisicakti ve outputta 2 tane print vericekti.
// ** go func haliyle calistirdigimizda sadece "hello from main" yazdirdi.Cunku go-routinei calistiricak bir vakti olamadi ; cunku bunu bekleyen bir mekanizma yok.
// ** tekrar calistirdigimizda 2 print veredebilir, vermeyedebilir. Suan UNDETERMINISTIC

// func main() {
// 	go func() {
// 		fmt.Println("Hello from Go-routine")
// 	}()

// 	fmt.Println("Hello from Main")
// }
