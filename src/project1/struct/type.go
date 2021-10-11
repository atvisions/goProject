package main

import "fmt"

// myInt 自定义类型
type myInt int

//结构体
type Persons struct {
	Name string
	Age  int
	Hobby []string
	map1 map[string]string

}

func main(){
	var a myInt =  20
	fmt.Printf("%v %T",a,a)
	//实例化
	var p Persons
		p.Name = "Cai Cai"
		p.Age  = 38
		p.Hobby = make([]string ,3,6)
		p.Hobby[0] = "打台球"
		p.Hobby[1] = "音乐"
		p.Hobby[2] = "钓鱼"

		p.map1 = make(map[string]string)
		p.map1["address"] = "上海"
		p.map1["telPhone"] = "13002175556"

	fmt.Println(p)
	fmt.Println(p.map1)
}
