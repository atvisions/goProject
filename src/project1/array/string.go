package main

import "fmt"

var str = "hello"

func main(){
	//定义一个byte类型的变量
	var scaleByte []byte
	//字符串先强制转换撑byte类型
	scaleByte = []byte(str)
	scaleByte[3] = 'o'
	str = string(scaleByte)
	fmt.Printf("%s",str)

}
