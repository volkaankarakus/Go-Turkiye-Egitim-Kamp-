// ** TAKIPCI ORNEGINDE TAKIPCI SAYISI 10'A ULASILINCA HATA VERSIN ISTIYORUM
// ** Go'daki errorlar aslinda bir interface tanimi.

package main

import (
	"errors"
)

type FlowUser struct {
	ID        int        `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Surname   string     `json:"surname,omitempty"`
	Followers []FlowUser `json:"followers,omitempty"`
}

// ** normal receiver
func (mainUser FlowUser) GetFollowerNumber() int {
	return len(mainUser.Followers)
}

// ** pointer receiver
func (mainUser *FlowUser) AddFollower(willAddUser FlowUser) error {
	if mainUser.GetFollowerNumber() == 4 {
		return errors.New("reached max user length (4)")
	}
	if mainUser.Followers == nil {
		mainUser.Followers = []FlowUser{}
	}
	mainUser.Followers = append(mainUser.Followers, willAddUser)
	return nil
}

// func main() {
// 	user := FlowUser{
// 		ID:      1,
// 		Name:    "volkan",
// 		Surname: "karakus",
// 		Followers: []FlowUser{
// 			{
// 				ID:      2,
// 				Name:    "ege",
// 				Surname: "okuyucu"},
// 			{
// 				ID:        3,
// 				Name:      "baki",
// 				Surname:   "can",
// 				Followers: []FlowUser{},
// 			}},
// 	}

// 	willAddUser := FlowUser{
// 		ID:      4,
// 		Name:    "onur",
// 		Surname: "bulut",
// 	}

// 	fmt.Println(user.GetFollowerNumber())
// 	if err := user.AddFollower(willAddUser); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(user.GetFollowerNumber())

// }
