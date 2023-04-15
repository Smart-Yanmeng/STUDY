package main

import "fmt"

// SayHello - 接口的定义
type SayHello interface {
	// 声明没有实现的方法
	sayHello()
}

// Chinese - 中国人
type Chinese struct {
}

// America - 美国人
type America struct {
}

// 实现接口的方法
func (ch Chinese) sayHello() {
	fmt.Println("你好")
}

func (am America) sayHello() {
	fmt.Println("Hello")
}

// 函数调用接口
func greet(s SayHello) {
	s.sayHello()
}

func main() {
	// 创建一个中国人
	c := Chinese{}
	greet(c)

	// 创建一个美国人
	a := America{}
	greet(a)
}
