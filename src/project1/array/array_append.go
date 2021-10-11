package main

import "fmt"

func main(){
	array := []int{1,2,3,4,5}
	for i:=6;i<10;i++ {
		array = append(array,i)

	}
	fmt.Printf("array:%v",array)
}
