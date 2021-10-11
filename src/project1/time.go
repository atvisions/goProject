package main

import (
	"fmt"
	"time"
)

var now = time.Now()
//var nows = time.Now().Unix()

func main(){
	//年
	var year = now.Year()
	//月
	var month = now.Month()
	//日
	var day = now.Day()

	//格式化
	timeStr := now.Format("2006/01/02 15:04:05")
	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(day)
	fmt.Printf("当前时间：%s",timeStr)
	//fmt.Println(now)
	time.Sleep(time.Second * 3)
}
