package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	byteStr , err := ioutil.ReadFile("src/project1/file/text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byteStr))
}
