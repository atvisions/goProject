package main

import (
	"fmt"
	"os"
)

func main(){
	err := os.MkdirAll("./001",006)
	if err != nil {
		fmt.Println(err)
		return
	}
}
