package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//存入数字
func putNum(intChan chan int) {
	for i:= 2;i < 100;i++{
		intChan <- i
	}
	close(intChan)
	wg.Done()
}
//存入素数
func primeNum(intChan chan int,primeChan chan int,exitChan chan bool)(){
	for num:= range intChan{
		var flag = true
		for i:=2;i<num;i++{
			if num%i == 0{
				flag = false
				break
			}
		}
		if flag{
			primeChan <- num
		}
	}
	//执行一次向exitChan写入一条数据
	exitChan <- true
	wg.Done()
}

//打印素数
func printNum(primeChan chan int){
	for v:= range primeChan{
		fmt.Println(v)
	}
	wg.Done()
}
func main(){
	//存放数字
	intChan := make(chan int,1000)
	//存放数字
	primeChan := make(chan int,1000)
	//标识素数存入结束
	exitChan := make(chan bool,16)
	//存放数字的协程
	wg.Add(1)
	go putNum(intChan)
	//存放素数的协程
	for i:=0 ;i<16;i++{
		wg.Add(1)
		go primeNum(intChan,primeChan,exitChan)
	}
	go func(){
		for i:=0;i<16;i++{
			<- exitChan
		}
		close(primeChan)
	}()
	//打印素数的协程
	wg.Add(1)
	go printNum(primeChan)
	wg.Wait()

}
