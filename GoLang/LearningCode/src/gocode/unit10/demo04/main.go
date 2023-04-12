package main

import "fmt"

type Student struct {
	Name string
}

// 传值函数
func func01(s Student) {
	fmt.Println(s.Name)
}

// 传指针函数
func func02(s *Student) {
	fmt.Println((*s).Name)
}

// 传值方法
func (s Student) method01() {
	fmt.Println(s)
}

// 传指针方法
func (s *Student) method02() {
	fmt.Println((*s).Name)
}

func main() {
	s := Student{"York"}

	// 调用传值函数
	func01(s)
	// func01(&s) -> 报错

	// 调用传指针函数
	func02(&s)
	// func02(s) -> 报错

	// 调用传值方法
	s.method01()
	// 使用传指针方式调用也没有报错，传递依旧是按照传值方式
	(&s).method01()

	// 调用传指针方法
	(&s).method02()
	// 使用传值方式调用也没有报错
	s.method02()
}
