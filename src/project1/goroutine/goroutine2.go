package main

import (
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func say(n int){
	num := 1000
	if n >= 0 {
		for i:= n * num ;i < (n+1) * num;i++{
			fmt.Printf("协程（%v）打印(%v)\n",n,i)
		}
	}


	defer wg1.Done()
}

func main(){
	start := time.Now().Unix()
	for i:=-1 ; i<= 5;i++ {
		wg1.Add(1)
		go say(i)
	}
	time.Sleep(time.Microsecond * 1000)
	wg1.Wait()
	fmt.Println("主进程结束...")
	end := time.Now().Unix()
	time := end - start
	fmt.Println(time)
}