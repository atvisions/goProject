package main

import "fmt"

type Student struct {
	Name string
	Age  int
	Sex  string
}

type Class struct {
	Title string
	Students []Student
}

func main(){
	student := []Student{
		{Name: "cc",Age: 42,Sex: "男"},
		{Name: "刘朝",Age: 32,Sex: "男"},
		{Name: "菜菜",Age: 22,Sex: "女"},
	}
	c := Class{
		Title: "001班",
		Students: make([]Student,0),
	}

	for i:= 0;i<len(student);i++{
		var s = student[i]
		c.Students = append(c.Students,s)
	}
	fmt.Println(c)
	//fmt.Printf("%v %v",c,len(person))
}
