package main

import "fmt"


func main()  {
	//创建channel
	ch := make(chan int,4)
	ch <- 10
	ch <- 28
	ch <- 37
	ch <- 45
	close(ch)  //关闭管道，适用range
	//fmt.Printf("地址：%v 容量：%v 长度：%v 值：%v",ch,cap(ch),len(ch),num)
	//for i:=0;i<cap(ch);i++{
	//	num := <- ch
	//	fmt.Println(num)
	//}
	<-ch //10取出了
	for v:= range ch{
		fmt.Println(v)
	}
}