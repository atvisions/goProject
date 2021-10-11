package main

import "fmt"

func sorts(arr []int) []int{
	for n:=0 ;n<len(arr);n++{
		for i:= 1;i< len(arr) - n;i++ {
			if arr[i] < arr[i - 1] {
				arr[i], arr[i - 1] = arr[i - 1], arr[i]
			}
		}

	}
	return arr
}

func main(){
	//arr := [...]int64 {5,-30,20,10,8}
	//arr[0] = 1
	//arr[2] = 5

	//遍历数组，打印key和value
	//for k,v:=range arr {
	//	fmt.Printf("key:%v Value：%v \n",k,v)
	//}
	//找出数组中数值最大的数
	//var max int64 = arr[0]
	//var index int = 0
	//for i:=0;i<len(arr);i++{
	//	if max < arr[i] {
	//		max = arr[i]
	//		index = i
	//	}
	//}
	//fmt.Printf("index = %v max = %v",index,max)
	//for k,v:=range arr {
	//
	//	//fmt.Printf("key:%v Value：%v \n",k,v)
	//}

	//fmt.Println(arr)

	//数组从小到重新排序
	array := [...]int {5,-30,20,10,8,854,756,-142,1,17}

	//sortSlice := sort(array[:]) //获取数组里的所有值
	sortSlice := sorts(array[0:2]) //获取数组里对应小标的值
	fmt.Printf("切片排序：%v \n",sortSlice)
	fmt.Printf("切片容量：%v \n",cap(sortSlice))
	fmt.Printf("切片长度：%v",len(sortSlice))
}
