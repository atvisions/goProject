package main

import (
	"fmt"
)

func main(){
	//str := "你好！china 2020"
	//
	//for _,v := range str{    //range 相当于rune类型
	//	fmt.Printf("原样字符：%c ascii：%v \n",v,v)
	//}
	//time.Sleep(time.Second * 3)

	str := "你好！golang"
	strRune := []rune(str)
	strRune[0] = '我'   //这里要使用单引号
	fmt.Println(string(strRune))


}

//中文需要使用range ，英文字符可以使用for i:=0;i<str;i++