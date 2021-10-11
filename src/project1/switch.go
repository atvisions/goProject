package main

import "fmt"

//func main(){
//	var value int = 4
//	switch value {
//	case 1:
//		fmt.Printf("%v太小",value)
//	case 2:
//		fmt.Printf("%v太大",value)
//	default:
//		fmt.Printf("%v这个数值ok",value)
//	}
//}

func main(){
	num := 1751
	switch  {
	case num < 100:
		fmt.Printf("%v数值太小",num)
	case num > 1000:
		fmt.Printf("%v数值太大",num)
	default:
		fmt.Printf("%v数值刚好",num)

	}
}