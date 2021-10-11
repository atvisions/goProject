package main

import (
	"fmt"
	"sync"
	"time"
)
var wg2 sync.WaitGroup
//写入数据
func fn1(ch chan int){
	for i:= 1;i <= 10;i++{
		ch <- i
		fmt.Printf("【写入】数据：%v成功！\n",i)
		time.Sleep(time.Microsecond * 10000)
	}
	wg2.Done()
	close(ch)
}
//写入数据
func fn2(ch chan int)  {
	for v:= range ch{
		fmt.Printf("【读取】数据：%v成功！\n",v)
		time.Sleep(time.Microsecond * 100000)
	}
	wg2.Done()
}

func main(){
	//创建管道
	ch := make(chan int,10)
	wg2.Add(1)
	go fn1(ch)
	wg2.Add(1)
	go fn2(ch)
	wg2.Wait()
}
