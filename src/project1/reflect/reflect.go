package main

import (
	"fmt"
	"reflect"
)
//reflect反射获取任意变量的类型

type Person struct {
	Name string
	Age  int
}

func reflectFn(x interface{}){
	//判断空接口类型
	v:= reflect.TypeOf(x)
	fmt.Printf("类型：%v  类型名称:%v  类型种类:%v \n",v,v.Name(),v.Kind())
}

func main()  {
	var myInt int
	var a = 10
	var b = 10.12
	var c = "welcome"
	d := make(map[string]interface{}) ;{
		d["name"] = "sss"
		d["age"] = 29
	}

	fmt.Println(d)

	reflectFn(a)
	reflectFn(b)
	reflectFn(c)
	reflectFn(d)
	reflectFn(myInt)
	reflectFn(Person{})

}
