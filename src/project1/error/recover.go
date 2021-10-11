package main

import "fmt"

func r1(x , y int) int {
	//抛出异常，延迟执行
	defer func(){
		err := recover()
		if err != nil{
			fmt.Printf("err:%v",err)
		}
	}()
	return x/y
}

func main(){
	fmt.Println(r1(10,0))
}