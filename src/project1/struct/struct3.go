package main

import "fmt"

type  Action struct {
	Name string

}

type Dog struct {
	age string
	Action
}

func (a Action) run(){
	fmt.Printf("%v今年在奔跑 \n",a.Name)
}

func (d Dog) gou(){
	fmt.Printf("%v今年%v岁了，他在汪汪的叫",d.Name,d.age)
}
func main(){
	var d = Dog{
		age: "2",
		Action:Action{
			Name: "丢丢",
		},
	}
	///d.run()
	d.gou()
}