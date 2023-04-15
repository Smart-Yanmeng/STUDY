package main

import (
	"fmt"
	"reflect"
)

// 定义学生结构体
type Student struct {
	Name string
	Age  int
}

func testReflect(i interface{}) {
	reType := reflect.TypeOf(i)
	fmt.Println(reType)

	reVal := reflect.ValueOf(i)
	fmt.Println(reVal)

	/* 获取变量的类别 */
	// reType
	k1 := reType.Kind()
	fmt.Println(k1)

	// reVal
	k2 := reVal.Kind()
	fmt.Println(k2)

	i2 := reVal.Interface()
	// 类型断言
	n, flag := i2.(Student)
	if flag {
		fmt.Printf("结构体的类型是：%T", n)
	}
}

func main() {
	// 对结构体类型进行反射
	// 定义结构体具体的实例
	stu := Student{
		Name: "York",
		Age:  18,
	}

	testReflect(stu)
}
