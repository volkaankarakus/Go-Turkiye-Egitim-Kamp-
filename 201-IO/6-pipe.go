package main

import (
	"fmt"
	"io"
	"time"
)

// ** bufio icin demistik ki, bazen memoryde bazı sizeda datayı bufferlayarak islem yapmak isteyebiliriz.
//
//	** bazen de memoryde hicbir buffer olmadan bir noktadan bir noktaya veri tasimak isteyebiliriz .
//
// ** İki farklı go-routineimiz olsun. Birisi okuma ve digeri yazma
// ** Bir boru düşünelim birisi bir ucundan yazıyor, digeri diger taraftan bakarak okuyor (arada buffer yok)
//
//	** bu ikisinin ortak haberlesmesini istiyorsak :
// func main() {
// 	pipeReader, pipeWriter := io.Pipe()
// 	done := make(chan struct{}) // bu channelin amacı sadece go-routinelerin islemlerini bitirmesini beklemesi icin

// 	go read(pipeReader, done)
// 	go write(pipeWriter)

// 	<-done
// }

func write(writer *io.PipeWriter) {
	// ** asagidaki kodlardan once elimizde karısık bir data olsun.
	// ** bu veriyi isleyip biryerlere yazıcaz
	/*
		INCOMING BULK DATA
	*/

	i := 0
	for {
		if i == 10 {
			writer.Close() // Writer'i kapattigimizda Reader daha fazla daha gelmeyecegini ögrenir.
			break
		}
		writer.Write([]byte(string(i))) // bütün sayilari bu hedefe yazdırıyoruz.
		i++
		time.Sleep(time.Millisecond * 100) // gorebilelim diye sleep attik.
	}
}

// ** okuma kısmındaki detay, veri okumadan da error donmus olabilir, biraz veri okuyup sonra error vermis de olabilir.
func read(reader io.Reader, done chan struct{}) {
	buffer := make([]byte, 1024)
	for {
		readed, error := reader.Read(buffer) // hep 1024'luk doldurabildigin kadar oku.
		if readed == 0 {
			if error == io.EOF { // EOF, End of file. Okuma yaptigim yer kapanmıs mı diye bakıyoruz.(streamin bitmesi)
				done <- struct{}{} // bittiyse main go-routine'e bittigini haber veriyoruz. Empty struct kullanilma sebebi memoryde gereksiz alan kaplamamak
				// bu yapıya SIGNAL denir.
				break
			}

			if error != nil {
				fmt.Println(error)
				done <- struct{}{}
				break
			}
		} else {
			// ** data dondugu halde error donmus olabilir
			fmt.Println(buffer[:readed]) // okudugun kadarını bana yazdır.

			if error == io.EOF || error != nil {
				fmt.Println(error)
				done <- struct{}{}
				break
			}
		}
	}
}
