package main

import "fmt"

//闭包

/*
全局变量
1、常驻内存
2、污染全局

局部变量
1、常驻内存
2、不污染全局

闭包
1、可以让一个变量常驻内存
2、可以让一个变量不污染全局

*/


func adder() func(y int) int{
	var i = 10
	return func(y int) int{
		i += y   //重新赋值i的值，后续再调用就调用了新的 i
		return i
	}
}
func main(){
	var fn = adder()
	fmt.Println(fn(15))
	fmt.Println(fn(15))
	fmt.Println(fn(15))
}
