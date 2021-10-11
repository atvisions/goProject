package main

import (
	"fmt"
	"sort"
)

func main(){
	sliceA := []string{"golang","php","java","nodejs"}
	sliceB := []int{10,21,3,41,55,6}
	sliceC := []float64{0.110,2.1,7.3,0.41,52.5,8.6}
	//字符串排序
	sort.Strings(sliceA)
	//int排序
	sort.Ints(sliceB)
	//浮点排序
	sort.Float64s(sliceC)

	fmt.Println(sliceA)
	fmt.Println(sliceB)
	fmt.Println(sliceC)
}
