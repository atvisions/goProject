package main

import (
	"os"
	"strconv"
)

func main(){
	file,err := os.OpenFile("src/project1/file/write.txt",os.O_CREATE|os.O_RDWR|os.O_WRONLY,0666)
	defer file.Close()
	if err != nil {
		return
	}
	//写入文件
	for i:= 0;i<10;i++ {
		file.WriteString("写入一句话 "+strconv.Itoa(i) +"\r\n")
	}


}
