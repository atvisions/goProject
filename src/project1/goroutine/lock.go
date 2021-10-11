package main

import (
	"fmt"
	"sync"
	"time"
)

//以下代码不加锁会导致读写数据丢失
var lock sync.Mutex //互斥锁
var wgs sync.WaitGroup


var m = make(map[int]int,0)

func test(num int){
	lock.Lock()
	var sum = 1
	for i:=1;i<=num;i++{
		sum *= num
	}
	m[num] = sum
	fmt.Printf("key=%v value=%v \n",num,sum)
	lock.Unlock()
	wgs.Done()
}
func main(){
	for i:=0;i<100;i++{
		wgs.Add(1)
		go test(i)
	}
	time.Sleep(time.Microsecond * 10000)
	wgs.Wait()
}