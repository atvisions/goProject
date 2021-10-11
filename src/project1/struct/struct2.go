package main

import "fmt"

type NewUser struct {
	Account  string
	Password string
	email    string
}

type UserInfo struct {
	address string
	sex     string
	hobby   []string
	NewUser
}


func main(){
	var uu UserInfo
	uu.NewUser.Account = "siYoYo"
	uu.address = "sh"
	uu.sex = "man"
	uu.hobby = make([]string,3)
	uu.hobby[0] = "taiQiu"
	uu.hobby[1] = "music"
	uu.sex = "nv"

	fmt.Printf("%#v \n",uu)

}


