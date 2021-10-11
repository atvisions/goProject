package main

import (
	T "demo/package/module" //T 包别名 可以简写module
	"fmt"
)

func main(){
	res := T.Num(19,50)
	fmt.Println(res)
}