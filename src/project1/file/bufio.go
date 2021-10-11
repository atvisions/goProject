package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main()  {
	file,err := os.Open("src/project1/time.go")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	if err != nil {
		fmt.Println(err)
		return
	}
	//逐行读取文件
	reader := bufio.NewReader(file)
	for {
		a, _, c := reader.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(string(a))
	}

}
