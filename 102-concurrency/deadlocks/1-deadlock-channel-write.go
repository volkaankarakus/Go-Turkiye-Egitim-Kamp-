package main


// ** UNBUFFERED bir channela hicbir dinleyen olmadigi sürece birşeyler yazmaya calisirsak ve tek bir goroutine varsa bir deadlock olusturur.
// func main() {
// 	unbufferedChannel := make(chan int)

// 	unbufferedChannel <- 1

// 	fmt.Println("hello unbuffered write deadlock")

// }

// OUTPUT : fatal error: all goroutines are asleep - deadlock!
