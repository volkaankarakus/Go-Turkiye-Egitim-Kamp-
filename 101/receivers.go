// ** GO'da 2 farkli function var
// ** Func ve Receiver
// ** receiverlar Darttaki factory gibi. Model uzerindeki islemler icin kullaniyoruz

// ** receiverlari da ikiye ayirabiliriz. Normal receiverlar ve pointer receiverlar

package main

type UserModel2 struct {
	ID        int          `json:"id,omitempty"`
	Name      string       `json:"name,omitempty"`
	Surname   string       `json:"surname,omitempty"`
	Followers []UserModel2 `json:"followers,omitempty"`
}

// ** Bir degisiklik yok pointera gerek yok. (normal receiver)
func (user UserModel2) FollowerNumber() int {
	return len(user.Followers)
}

// ** POINTERDAN ALMADIGIMIZ ICIN ISE YARAMAZ. ASAGIDAKI GIBI POINTERDAN ALMAMIZ GEREKIYOR.
// ** STRUCTLARA YAPILACAK DEGISIKLIKLER ICIN DE POINTER KULLAN (suan normal receiver)
func (mainUser UserModel2) AddFollower(willAddUser UserModel2) {
	if mainUser.Followers == nil {
		mainUser.Followers = []UserModel2{}
	}
	mainUser.Followers = append(mainUser.Followers)

}

// ** asagidaki islemi yaptigimizda son uzunluk yine ayni gelir. Cunku burada da degisiklikleri pointerdan takip etmemiz gerekiyor
// FollowerNumber()
// AddFollower()
// FollowerNumber()

// ** STRUCTTA DEGISIKLIK YAPICAKSAK POINTER RECEIVER TANIMLAMAMIZ GEREKIYOR KI ADRESTEN ALSIN
// ** (pointer receiver)
func (mainUser *UserModel2) AddFollowerFromPointer(willAddUser UserModel2) {
	if mainUser.Followers == nil {
		mainUser.Followers = []UserModel2{}
	}
	mainUser.Followers = append(mainUser.Followers, willAddUser)
}

// func main() {
  
// }
