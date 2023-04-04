package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// ** Bir stream okundugu zaman, 2. kez bunu okuyamayiz cunku stream bitmistir.
// ** Oyle birsey istiyoruz ki streami duplicate edelim. -> io.TeeReader
func main() {
	stringReader := strings.NewReader("test string")

	// ** duplicates the stream
	teeReader := io.TeeReader(stringReader, os.Stdout) // teeReader reader ve writer alıyor, reader donuyor.
	fmt.Println("we will start")

	readedBytes, _ := io.ReadAll(teeReader) // readAll, streamdeki butun datayi okumayı saglar.
	fmt.Println("\n --- readed string ---")
	fmt.Println(string(readedBytes))

	/*
		OUTPUT :
			we will start
			test string
			--- readed string ---
			test string
	*/
	// ** BUNU NEREDE KULLANABILIRIZ ?
	// ** Bize bir Http requesti geliyor. Http body'sini 1 kere okuduktan sonra ikinciye okuyamayiz.
	// ** 	Biz bunu biryere kaydedip uzerinde degisiklikler yapmak istiyorsak TeeReader ile streami duplicate edebiliriz.
}
