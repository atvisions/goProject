package main

import "fmt"

func main(){
	//切片和数组的区别就是 数组不能扩容
	//定义切片
	//var slice1 []int64

	//或使用 make() 函数来创建切片:
	//slice2 := make([]int64, 10 ,20)

	//切片初始化
	//slice1 = []int64 {1,2,3,4,5,6}
	//slice2 = []int64 {1,2,3,4,5,6}
	//fmt.Println(slice2)

	//golang分配内存有一个make函数，该函数第一个参数是类型，第二个参数是分配的空间，第三个参数是预留分配空间，前两个参数都很好理解，
	//对于第三个参数，例如a:=make([]int, 5, 10)， len(a)输出结果是5，cap(a)输出结果是10，对a[4]进行赋值发现是可以得，
	//但对a[5]进行赋值发现报错了，那么这个预留分配的空间要怎么使用呢？预留的空间需要重新切片才可以使用。
	a := make([]int, 5, 10)

	//切片扩容 append()
	a = append(a,18)
	fmt.Println(a)

	fmt.Printf("%d, %d\n", len(a), cap(a))
	fmt.Println(a)
	b := a[:cap(a)]
	fmt.Println(b)
}
