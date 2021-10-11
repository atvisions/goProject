package main

import (
	"fmt"
	"io"
	"os"
)

func main(){

	//打开文件
	file,err:= os.Open("src/project1/reflect/reflect.go")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	//读取文件内容
	var strSlice []byte
	var tempSlice = make([]byte,256)
	for  {
		n,err := file.Read(tempSlice)
		if err != nil && err != io.EOF {
			fmt.Printf("error:%v \n",err)
		}
		if n == 0 {
			break
		}
		fmt.Println(n)
		strSlice = append(strSlice,tempSlice[:n]...)   //注意这里的写法
	}
	fmt.Printf("文件内容：%v",string(strSlice))

}
