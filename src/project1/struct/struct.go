package main

import "fmt"
type Info struct {
	Name string
	Sex  int64
	Age  int64
	Add  string
}

type User struct {
	NickName string
	Info //结构体继承  结构体嵌套
}

type Admin struct {
	AdminName string
	Info
}

func main() {
	// 声明方式1 user1 为 User 类型
	user1 := User{
		"飞天大老鼠",
		Info{
			Sex: 1,
			Name: "刘朝",
			Age: 18,
			Add: "广州市白云区"},

	}
	// 声明方式2 user2 为 User 类型
	var u User

	u.Name = "李四"
	u.NickName = "菜菜"
	u.Sex = 0
	u.Add = "上海市长宁区"
	u.Age = 41

	var au Admin
	au.AdminName = "admin"
	au.Name = "cc"

	fmt.Println(user1)
	fmt.Println(au)
}
