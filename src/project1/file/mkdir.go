package main

import (
	"fmt"
	"os"
)

func main(){
	err := os.Mkdir("./src/project1/mkdir",0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}
