package main

import "fmt"

func main(){

	var value [...]string


	var data1 = []string {"php","beego","golang"}
	for i:= 0;i<len(data1);i++{
		value[i] = data1[i]
	}
	fmt.Printf("%v",value)
}
