package main

import "fmt"

type Student struct {
	Name string
}

// 定义一个方法
func (s Student) method01() {
	fmt.Println("方法输出：", s.Name)
}

// 定义一个函数
func func01(s Student) {
	fmt.Println("函数输出：", s.Name)
}

func main() {
	// 调用函数
	var s Student = Student{"York"}
	func01(s)

	// 调用方法
	s.method01()
}
