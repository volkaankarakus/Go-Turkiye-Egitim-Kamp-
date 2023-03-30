package main

import (

)

// ** Normalde asagidaki fonksiyonun 1'den 10'a kadar rastgele cikti vermesini bekleriz.
//    ** ama buradaki output : 10 10 10 10 10 10 10 10 10 10
//    ** Bunun yasanmamasi icin her dongude referans bir copy olusturup (go func sonunda), buraya(go func basına) onu gecmeliydik.
// ** Burada SCOPING ve REFERENCE kavramı var.

// func main() {
// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			fmt.Println(i)
// 		}()
// 	}
// 	time.Sleep(time.Second * 3)
// }

// ** Bu haliyle calistirdigimizda : 8 9 1 6 7 3 0 5 4 2 verdi.
// func main() {
// 	for i := 0; i < 10; i++ {
// 		go func(val int) {
// 			fmt.Println(val)
// 		}(i)
// 	}
// 	time.Sleep(time.Second * 3)
// }
