package main

import "fmt"

type Address struct {
	Name string
	Add  string
}

func main(){
	var userInfo = make(map[string]interface{})
	userInfo["name"] = "张三"
	userInfo["age"] = 28
	userInfo["address"] = Address{
		Name: "王思嘉",
		Add: "长宁区清池路",
	}
	userInfo["hobby"] = []string{"台球","音乐","钓鱼"}

	//fmt.Println(userInfo["hobby"][0]) //invalid operation: userInfo["hobby"][0] (type interface {} does not support indexing) 接口类型，不能找到索引值

	//hobby类型断言
	var u = userInfo["hobby"]
	if v,ok := u.([]string);ok {
		hobby := v
		fmt.Println(hobby[0])
	}
	//address类型断言
	var a = userInfo["address"]
	if v,ok := a.(Address);ok {
		address := v
		address.Name = "李四"
		fmt.Println(address.Name)
	}
	//fmt.Println(userInfo)
}
