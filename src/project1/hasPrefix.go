package main

import (
	"fmt"
	"strings"
	"time"
)

func sub(a int, b int) (add int ,reduce int){
	add = a + b
	reduce = a - b
	return add,reduce
}

func main(){

	//a := 678
	//b := true
	//c := 1.345
	//d := "jon"
	//
	////fmt.Println(sub(13,6))
	//fmt.Printf("a = %d b = %t c = %f d = %s \n",a,b,c,d)
	str := "http://www.baidu.com"
	ips := "8.8.8.8;114.114.114.114"

	if strings.HasPrefix(str,"http"){
		fmt.Print("has http\n")
	}

	if strings.HasSuffix(str,"baidu.com"){
		fmt.Printf("has baidu.com\n")
	}
	lens := len(str)
	fmt.Printf("%d\n",lens)

	IParry := strings.Split(ips,";")
	fmt.Printf("ip1:%s\n", IParry[0])
	fmt.Printf("ip2:%s\n", IParry[1])


	time.Sleep(time.Second * 10)
}
