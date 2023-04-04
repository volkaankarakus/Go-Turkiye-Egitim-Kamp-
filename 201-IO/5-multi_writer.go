package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	txtFile1, _ := os.Create("txtFile1.txt")
// 	txtFile2, _ := os.Create("txtFile2.txt")
// 	multiWriter := io.MultiWriter(os.Stdout, txtFile1, txtFile2)

// 	howManyBytesInt, error := multiWriter.Write([]byte("multi writer example"))
// 	if error != nil {
// 		fmt.Println(error)
// 	}

// 	txtFile1.Close()
// 	txtFile2.Close()
// 	fmt.Println(howManyBytesInt) // 20
// }
