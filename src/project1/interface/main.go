package main

import "fmt"

// Usb 定义一个usb的接口
//接口是一组方法的集合
type Usb interface {
	Start()
	Stop()
}

// Phone 定义一个手机的结构体，如果接口里有方法的话需要通过结构体或者自定义方法类型实现这个接口
type Phone struct {
	Name string
}

// Camera 结构体内可以什么都不写
type Comera struct {

}

// Start 定义一个结构体的方法，p phone 接受者
func (p Phone) Start() {
	fmt.Printf("%v启动 \n",p.Name)
}
func (p Phone) Stop() {
	fmt.Printf("%v关闭 \n",p.Name)
}

// Start 相机，方法必须完整对应接口，否则报错
func (c Comera) Start(){
	fmt.Println("相机启动")
}
func (c Comera) Stop(){
	fmt.Println("相机关闭")
}

func (c Comera) Run(){
	fmt.Println("run")
}

func main(){
	//实例化
	var p = Phone{
		Name: "小米手机",
	}

	var pp Usb  //接口就是一个数据类型
	pp = p
	pp.Start()
	pp.Stop()

	//相机实例化
	c := Comera{}
	//相机实现了接口
	var c1 Usb = c
	//相机实现方法
	c1.Start()
	c1.Stop()
	//c1.Run() //c1.Run undefined (type Usb has no field or method Run) 接口中没有run方法
	c.Run()

}