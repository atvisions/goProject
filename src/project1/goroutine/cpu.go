package main

import (
	"fmt"
	"runtime"
)

func main()  {
	//获取当前计算机cpu数量
	var cpuNum = runtime.NumCPU()
	fmt.Println(cpuNum)
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
