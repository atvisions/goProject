package main

import "fmt"


//匿名函数
func main(){
	func(){
		fmt.Println("text.....")
	}()  //结尾括号可以执行

	var fn = func (x,y int) int{
		return x * y
	}

	fmt.Println(fn(8,9))

	//匿名自执行函数也可以接收参数
	var t = func (x,y int) int{
		return x + y
	}(18,98)

	fmt.Println(t)
}