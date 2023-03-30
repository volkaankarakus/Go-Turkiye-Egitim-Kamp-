// ** GO'da exception yok.
// ** Ama uygulamayi kiran, uygulamanın calismasini saglayan PANIC fonksyionu var.

// ** Mesela bir HTTP server ayaga kaldiriyoruz, 80 portu daha onceden kullaniliyorsa panic atar.
// ** PANICleri attiktan sonra DEFER ile RECOVER kullanarak panic yakaliyoruz.

// ** Recover : Fonksiyon cagrisi icinde bir panic olustugu durumda bize bunlari bildirir ve burada islem yaptirabiliriz.(print gibi)
package main

import (
	"fmt"
	"time"
)

// func main() {
// 	timer(time.After(5 * time.Second))
// }

func timer(c <-chan time.Time, message ...string) {
	defer fmt.Println("bir sonraki asamaya geciliyor")
	defer fmt.Println("timer bitti")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	for {
		select {
		case <-c:
			return
		default:
			fmt.Println(time.Now(), message)
			time.Sleep(1 * time.Second)

			if time.Now().Day() == 30 {
				panic("beklenmeyen bir hata olustu")
			}
		}
	}
}
