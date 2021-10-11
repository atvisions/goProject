package main

import "fmt"

// MyPrint 类型断言
func MyPrint(X interface{}) {
	if v, ok := X.(string); ok {
		fmt.Printf("值是：%v 类型是：%T \n", v, v)
	}
}

func MyPrint2(X interface{}) {
	switch X.(type) {
	case string:
		fmt.Printf("%v是%T类型\n", X, X)
	case bool:
		fmt.Printf("%v是%T类型\n", X, X)
	case int:
		fmt.Printf("%v是%T类型\n", X, X)
	default:
		fmt.Println("未知类型")
	}
}

func main() {
	MyPrint("ssssss")
	MyPrint2(345)
	MyPrint2(true)
	MyPrint2("我爱你")
	MyPrint2(3.145)
}
