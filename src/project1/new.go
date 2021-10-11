package main

import "fmt"

func main(){
	//var a *int
	a := new(int)
	*a = 10
	b := *a

	fmt.Println(b)
}
