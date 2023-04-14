package main

import (
	"fmt"
	"reflect"
)

// 利用一个函数，函数的参数定义为空接口
func testReflect(i interface{}) {
	// 调用 ValueOf 函数，返回 reflect.Value 类型数据
	reVal := reflect.ValueOf(i)

	// 通过 SetInt() 来改变值
	reVal.Elem().SetInt(40)
}

func main() {
	// 对基本数据类型进行反射
	// 定义一个基本数据类型
	var num int = 100
	testReflect(&num)

	fmt.Println(num)
}
