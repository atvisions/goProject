package main

import "fmt"

func main(){
	var array = [...][2]string {
		{"上海","北京"},
		{"天津","成都"},
		{},

	}

	array[2][0] = "美国"
	fmt.Println(array)
}
