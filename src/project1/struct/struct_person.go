package main

import "fmt"

type Person struct{
	Name string
	age  int
	sex  string
}

func (p Person) personInfo(){
	fmt.Printf("姓名：%v 年龄：%v  性别:%v \n",p.Name,p.age,p.sex)
}

func (p *Person) setInfo(n string,a int){
	p.Name = n
	p.age = a
}

func main(){
	var p1 = Person{
		Name: "菜菜",
		age: 18,
		sex: "男",
	}
	var p2 = Person{
		Name: "cc",
		age: 18,
		sex: "女",
	}

	p1.personInfo()
	p2.personInfo()
	//如果setinfo 不使用指针引用这里就不会改变，显示还是菜菜
	p1.setInfo("caicai",28)
	p1.personInfo()
}
