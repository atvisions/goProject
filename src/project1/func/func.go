package main

import "fmt"

//自定义类型
//type calc func(int, int) int
//func add(x, y int) int {
//	return x + y
//}
//
//func sub(x, y int) int {
//	return x - y
//}
//
//func main() {
//	//变量指定类型
//	var a calc
//	var s calc
//	//变量可以使用和指定类型相同的函数
//	a = add
//	s = sub
//	fmt.Printf("%T %v %v", a, a(20, 30), s(70, 30))
//}

//将一个函数作为另一个函数的参数
//func add(x, y int) int {
//	return x + y
//}
//
//type calcType func(int,int)int
//
//func calc (x ,y int,cb calcType) int {
//	return cb(x,y)
//}
//func main(){
//	sum := calc(100,30,add)
//	fmt.Println(sum)
//}

//函数作为返回值
func add(x, y int) int {
	return x + y
}
func newSub(x, y int) int {
	return x - y
}

//定义一个方法类型
type calcType func(int,int)int


func do (o string) calcType{
	switch o {
	case "+":
		return add
	case "-":
		return newSub
	case "*":
		return func(x , y int) int {
			return x * y
		}
	default:
		return nil
	}
}

func main(){
	a := do("+")
	s := do("-")
	fmt.Println(a(50,80))
	fmt.Println(s(50,80))
}



