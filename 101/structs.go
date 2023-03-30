package main

// ** STRUCTS
// ** Struct icerisinde Darttaki classlara göre farkı : TAG'ler
// ** Tag'ler genelde refrection operasyonlarında kullanılıyor.
// ** Bir structi reflection ile gezdigimizde tagdeki degerine bakarak operasyon yapabiliyoruz.
// ** tagleri yazmazsak default olarak field isimlerini atar.

// ** field nil'se ve jsonda bunu basmak istemiyorsak OMITEMPTY kullanıyoruz.(readedByteArray'de yazmaz ama marshalsiz printlemede yazar)
import (
	// "encoding/json"
	// "fmt"
)

type UserModel struct {
	Id                                       int         `json:"id,omitempty"`
	Name                                     string      `json:"name,omitempty"`
	NameFieldIsNillAndWDontWanttoPrintInJson string      `json:"nameFieldIsNillAndWDontWanttoPrintInJson,omitempty"`
	Surname                                  string      `json:"surname,omitempty"`
	Followers                                []UserModel `json:"followers,omitempty"`
}

// func main() {
// 	// ** User1
// 	user1 := UserModel{
// 		Id:        1,
// 		Name:      "Volkan",
// 		Surname:   "Karakus",
// 		Followers: nil,
// 	}

// 	// ** User2
// 	user2 := UserModel{
// 		Id:      2,
// 		Name:    "Gungor",
// 		Surname: "Uludag",
// 		Followers: []UserModel{
// 			{
// 				Id:        3,
// 				Name:      "Kagan",
// 				Surname:   "Sahbaz",
// 				Followers: []UserModel{user1},
// 			},
// 			{
// 				Id:        4,
// 				Name:      "Emre",
// 				Surname:   "Aydin",
// 				Followers: nil,
// 			},
// 		},
// 	}

// 	// ** Marshal
// 	readedByteArrayUser1, marshalError := json.Marshal(user1)
// 	if marshalError != nil {
// 		fmt.Println(marshalError.Error())
// 		return
// 	}
// 	fmt.Println(string(readedByteArrayUser1)) // {"id":1,"name":"Volkan","surname":"Karakus"}
// 	fmt.Println(user1)                        // {1 Volkan Karakus []}

// 	readedByteArrayUser2, marshallErrorUser2 := json.Marshal(user2)
// 	if marshallErrorUser2 != nil {
// 		fmt.Println(marshalError.Error())
// 		return
// 	}
// 	fmt.Println(string(readedByteArrayUser2))
// 	fmt.Println(user2)

// }
