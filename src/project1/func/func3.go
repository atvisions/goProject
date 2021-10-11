package main

import "fmt"

//递归函数,自己调用自己

func fn1 (n int){
	if n > 0 {
		fmt.Println(n)
		n--
		fn1(n)
	}
}

//递归实现 1- 100 数字的相加
func fn2(n int) int{
	if n > 1{
		return n + fn2(n - 1)
	}else{
		return 1
	}

}

func main(){
	fn1(8)
	fmt.Println(fn2(365))
}
