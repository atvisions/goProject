package main

import (
	"encoding/json"
	"fmt"
)

type Student2 struct {
	Id  	int  	`json:"id"`
	Name 	string 	`json:"name"`
	Sex 	string	`json:"sex"`
	studentInfo2

}
type studentInfo2 struct {
	Hobby []string `json:"hobby"`
}

func main(){
	var str = `{"id":12,"name":"cc","sex":"man","hobby":["台球","音乐","钓鱼"]}`
	var s1 Student2
	err := json.Unmarshal([]byte(str),&s1)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v",s1)
}
