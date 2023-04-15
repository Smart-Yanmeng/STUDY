package main

import (
	"fmt"
	"reflect"
)

// 利用一个函数，函数的参数定义为空接口
func testReflect(i interface{}) {
	// 调用 TypeOf 函数，返回 reflect.Type 类型数据
	reType := reflect.TypeOf(i)
	fmt.Println(reType)

	// 调用 ValueOf 函数，返回 reflect.Value 类型数据
	reVal := reflect.ValueOf(i)
	fmt.Println(reVal)

	// 如果要获取 reVal 的数值，要调用 Int() 方法，返回 v 持有的有符号整数
	num := reVal.Int() + 80
	fmt.Println(num)

	// reVal 转成空接口
	i2 := reVal.Interface()
	// 类型断言
	n := i2.(int)
	n2 := n + 30

	fmt.Println(n2)
}

func main() {
	// 对基本数据类型进行反射
	// 定义一个基本数据类型
	var num int = 100
	testReflect(num)
}
