package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// func main() {
// 	//readFile()
// 	//readFileBufio()
// 	readFileSeek()
// }

func readFile() {
	readTest, error := os.ReadFile("testCreateFileAndWrite.txt")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(string(readTest))
	}
}

// ** IO islemleri yaparken bazı verileri bufferlamak isteyebiliriz.
// ** Bellekte belirli sayıda veriyi tutmak isteyebiliriz. Dosyadaki bütün veriyi memorye almadan okuyabiliriz : BUFIO

// ** IO'daki en onemli 2 konsept io.Reader ve io.Writer interfaceleri. Bu cok fazla karsımıza cıkıyor.
// ** NewReader'in icine gidersek, io.Reader alir. Bunun da icine gidersek :
// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }
// ** Bu interfaceler bize aslında soyutlama saglıyor.
// ** GO'da bu durum DUCK TYPING INTERFACE olarak gecer.
// ** DUCK TYPING INTERFACE'da extends, implements gibi kavramlar yok.
// 	**  Yukarıdaki ornekteki gibi Read aynı parametrelere sahipse otomatik olarak bu interface'i implement etmis demektir.
// **  (rd io.Reader) parametresi alan herhangi bir yere bu struct'i gönderebiliriz.(asagidaki gibi)
// func NewReader(rd io.Reader) *Reader{
// 	return NewReaderSize(rd,defaultBufSize)
// }

// ** readFileBufio'da da butun dosyayi okuduk. Parca parca olan daha sonraki fonksiyonda.
func readFileBufio() {
	readTest, error := os.Open("testCreateFileAndWrite.txt") // os.Open dosyanin pointerina erisir.
	if error != nil {
		fmt.Println(error)
	} else {
		bufReader := bufio.NewReader(readTest)
		io.Copy(os.Stdout, bufReader)
	}
}

// ** GB'larca datamız olsun ve biz bunun bir kısmını okumak istiyorsak :
func readFileSeek() {
	file, _ := os.Open("testCreateFileAndWrite.txt")
	file.Seek(3, 1)             // (offset int64, whence int)
	readByte := make([]byte, 4) // 3. offsettten sonra 4byte'lık okuma yap.
	file.Read(readByte)

	fmt.Println(string(readByte))
	// OUTPUT : 4567
}