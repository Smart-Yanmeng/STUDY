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

	n := val.Elem().NumField()
	fmt.Println(n)

	// 修改字段的值
	val.Elem().Field(0).SetString("张三")
}

func main() {
	s := Student{
		Name: "York",
		Age:  18,
	}

	// 调用 TestStudentStruct
	TestStudentStruct(&s)
	fmt.Println(s)
}
