package main

import "fmt"

//空接口类型可以是任意数据类型

func show(i interface{}){
	fmt.Printf("%v %T",i,i)
}

func main()  {
	var data =  []interface{}{111,"32323",true}
	fmt.Println(data)

	show("i love you")
}
