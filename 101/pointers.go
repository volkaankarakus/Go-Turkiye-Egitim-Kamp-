package main

import (
	"strings"
)

// ** POINTERS

// ** String'in zero valuesu boş string ” oldugu icin buna "nil" atayamayiz.
// ** nil atayabilecegimiz bir string istiyorsak bunu reference type'a cevirmemiz gerekiyor :
type String *string // artık stringi degil stringi point eden adresi tutar.

// func main() {
// 	var referenceString String
// 	var normalString string

// 	// fmt.Println(referenc eString) // OUTPUT  : <nil>
// 	// fmt.Println(normalString)    // OUTPUT :

// 	// ** bunlara bir degisken atayalim
// 	normalString = "Go turkiye"
// 	referenceString = &normalString

// 	fmt.Println(referenceString)  // OUTPUT  : 0x1400010c020
// 	fmt.Println(normalString)     // OUTPUT  : Go turkiye
// 	fmt.Println(*referenceString) // OUTPUT  : Go turkiye

// 	changeString(referenceString)
// 	fmt.Println(*referenceString) // OUTPUT : GO turkiye

// }

// ** String üzerinde herhangi bir degisiklik yapıcaksak bunu pointerindan yapmamız gerekir.
// ** degisiklik yapmak icin normal string uzerinde yaparsak bunu yaptıgımız yerde yapar ve öyle kalır.
// ** ama degisikligi adreste yaparsak(pointerında) kalıcı olarak kalır.
func changeString(s String) {
	*s = strings.Replace(*s, "Go", "GO", 1)
}
