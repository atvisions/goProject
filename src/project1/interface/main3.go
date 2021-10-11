package main

import "fmt"

// Usber Usb 定义一个接口
type Usber interface {
	start()
	stop()
}

// Phones 定义一个手机的结构体
type Phones struct {
	Name string
}
//定义一个相机的结构体
type Comeras struct {
}

//定义一个电脑的结构体
type Computer struct {

}
//定义手机的方法，值接收者
func (p Phones) start(){
	fmt.Printf("%v启动 \n",p.Name)
}
func (p Phones) stop(){
	fmt.Printf("%v关闭 \n",p.Name)
}

//定义相机的方法，指针接收者
func (c *Comeras) start(){
	fmt.Println("相机启动")
}
func (c *Comeras) stop(){
	fmt.Println("相继关闭")
}
//电脑的方法,顶一个work，需要传入usb的接口对象
func (c Computer) work(usb Usber){
	usb.start()
	usb.stop()
}

func main(){
	//实例化
	var computer = Computer{}
	var phones = Phones{
		Name: "小米",
	}
	var comers1 = &Comeras{}
	computer.work(phones)
	computer.work(comers1)

}


