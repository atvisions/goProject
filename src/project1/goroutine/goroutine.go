package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup

func Say1(s string){
	for i:=0;i<10;i++{
		time.Sleep(time.Microsecond * 1000)
		fmt.Printf("say1() %v-%v \n",s,i)
	}
	wg.Done() //计数器减1
}

func Say2(s string){
	for i:=0;i<10;i++{
		time.Sleep(time.Microsecond * 1000)
		fmt.Printf("say2() %v-%v \n",s,i)
	}
	wg.Done() //计数器减1
}

func main(){
	wg.Add(1) //计数器加1
	go Say1("world")  //开启一个协程
	wg.Add(1) //计数器加1
	go Say2("world")  //开启一个协程
	wg.Wait() //等待协程执行完毕
	fmt.Println("主线程退出...")
}

//执行以上代码，你会看到输出的 hello 和 world 是没有固定先后顺序。因为它们是两个 goroutine 在执行：