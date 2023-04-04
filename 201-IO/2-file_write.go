package main

import (
	"fmt"
	"os"
)

// func main() {
// 	//writeOs()
// 	createAndWriteFile()
// }

func writeOs() {
	// ** WriteFile :
	// 	** ilk parametre : dosya adi
	//  ** ikinci parametre : dosyanin icerisine yazilacak byteArray
	// 	** ucuncu parametre : bu dosyayi hangi yetkiyle acmak istiyoruz. ModePerm 0777 (Unix permission bits)
	err := os.WriteFile("testFile.txt", []byte("test"), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

func createAndWriteFile() {
	file, error := os.Create("testCreateFileAndWrite.txt")
	if error != nil {
		fmt.Println(error)
	}

	file.Write([]byte("1\n"))
	file.Write([]byte("2\n"))
	file.Write([]byte("3\n"))
	file.Write([]byte("4\n"))
	file.Write([]byte("5\n"))
	file.Close()
}
