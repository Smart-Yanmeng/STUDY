package main

import (
	"fmt"
	"reflect"
)

// Student - 定义学生结构体
type Student struct {
	Name string
	Age  int
}

func testReflect(i interface{}) {
	reType := reflect.TypeOf(i)
	fmt.Println(reType)

	reVal := reflect.ValueOf(i)
	fmt.Println(reVal)

	i2 := reVal.Interface()
	n, flag := i2.(Student)
	if flag {
		fmt.Println("学生的名字是：", n.Name)
		fmt.Println("学生的年龄是：", n.Age)
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
