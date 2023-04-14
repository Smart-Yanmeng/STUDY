package main

import (
	"fmt"
	"reflect"
)

// Student - 定义一个结构体
type Student struct {
	Name string
	Age  int
}

// Print - 给结构体绑定方法
func (s Student) Print() {
	fmt.Println("调用了 Print() 方法")
	fmt.Println("学生的名字是：", s.Name)
}

func (s Student) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Student) Set(name string, age int) {
	s.Name = name
	s.Age = age
}

// TestStudentStruct - 定义函数操作结构体进行反射操作
func TestStudentStruct(a interface{}) {
	// 将 a 转成 reflect.Value 类型
	val := reflect.ValueOf(a)
	fmt.Println(val)

	// 通过 reflect.Value 类型操作结构体内部的字段
	n1 := val.NumField()
	fmt.Println(n1)

	// 通过遍历获取具体的字段
	for i := 0; i < n1; i++ {
		fmt.Printf("第%d个字段的值是：%v\n", i, val.Field(i))
	}

	// 通过 reflect.Value 类型操作结构体内部的方法
	n2 := val.NumMethod()
	fmt.Println(n2)

	// 调用 cPrint() 方法
	// 调用方法，方法首字母必须大写
	val.Method(1).Call(nil)

	// 调用 GetSum 方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	result := val.Method(0).Call(params)

	fmt.Println("result:", result[0].Int())
}

func main() {
	s := Student{
		Name: "York",
		Age:  18,
	}

	// 调用 TestStudentStruct
	TestStudentStruct(s)
}
