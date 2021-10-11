package main

import (
	"fmt"
	"time"
)
//在某种场景下需要同时从多个管道获得数据，可以用select多路复用，不需要关闭管道，close()
func main(){
	intChan := make(chan int ,10)
	for i:=0;i<10;i++{
		intChan <- i
	}
	//close(intChan)

	stringChan := make(chan string,10)
	for i:=0;i<10;i++{
		stringChan <- "hello" + fmt.Sprintf("%d", + i)
	}
	for {
		select {
		case v:= <- intChan:
			fmt.Printf("从intChan读取%v成功！\n",v)
			time.Sleep(time.Microsecond * 1000000)
		case v:= <- stringChan:
			fmt.Printf("从stringChan读取%v\n",v)
		default:
			fmt.Println("数据读取完毕")
			return
		}

	}
}
