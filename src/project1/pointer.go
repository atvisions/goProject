package main

import "fmt"

func main(){
	var a int64 = 10
	var ip *int64
	var ptr *int
	ip = &a /* 指针变量的存储地址 */
	fmt.Println(ip) /* 指针变量的存储地址 */
	fmt.Println(&a) /* 指针变量的存储地址 */
	fmt.Println(ip) /* 指针变量的存储地址 */
	fmt.Println(*ip) /* 使用指针访问值 */

	if ptr != nil {
		fmt.Println(ptr)
	}
	fmt.Println("ptr是空指针")

}
