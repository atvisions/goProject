package main

import "fmt"

type Ainterface interface {
	SetName(string)
}

type Binterface interface {
	GetName()string
}

//接口嵌套

type Animals interface {
	Ainterface
	Binterface
}

// Dogs 定义一个dog结构体
type Dogs struct {
	Name string
}

// SetName 定义方法
func (d *Dogs) SetName(name string){
	d.Name = name
}

func (d *Dogs) GetName()string{
	return d.Name
}

func main(){
	var d = &Dogs{
		Name: "丢丢",
	}
	var d1 Animals = d
	d1.SetName("大黄")
	fmt.Println(d1.GetName())
}