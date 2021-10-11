package main

import "fmt"

// Animal 定义一个接口
type Animal interface {
	SetName(string)
	GetName()string
}

// Dog 定义一个dog的结构体
type Dog struct {
	Name string
}

// SetName 方法
func (d *Dog) SetName(name string){
	d.Name = name
}

func (d *Dog) GetName()string{
	return d.Name
}

func main(){
	//实例化
	d := &Dog{Name: "小黑"}
	//实现接口
	var d1 Animal = d
	fmt.Println(d1.GetName())
	d1.SetName("丢丢")
	fmt.Println(d1.GetName())
}