package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id  	int  	`json:"id"`
	Name 	string 	`json:"name"`
	Sex 	string	`json:"sex"`
	studentInfo

}
type studentInfo struct {
	Hobby []string `json:"hobby"`
}

func main(){

	var s1 = Student{
		Id: 1,
		Name: "cc",
		Sex: "男",
	}
	s1.Hobby = make([]string,3)
	s1.Hobby = []string{"台球","音乐","钓鱼"}
	jsonByte,_ :=  json.Marshal(s1)
	strJson := string(jsonByte)
	fmt.Println(strJson) //{"Id":38,"Name":"caicai","Sex":"man"}   首字母是大写的，改成小写需要用到结构体标签 `json:"id"`
}
