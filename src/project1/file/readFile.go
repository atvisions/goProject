package main

import (
	"fmt"
	"io/ioutil"
)
//复制文件

func copy(srcFile string,dstFile string)(err error){
	str ,err := ioutil.ReadFile(srcFile)
	if err != nil{
		return err
	}
	errs := ioutil.WriteFile(dstFile,str,006)
	if errs != nil{
		return errs
	}
	return nil
}

func main(){
	srcFile := "src/project1/file/text.txt"
	dstFile := "d:/text1.txt"

	err := copy(srcFile,dstFile)

	if err != nil{
		fmt.Println(err)
		return //没有这个return 程序不会结束
	}
	fmt.Println("复制文件成功")
}
