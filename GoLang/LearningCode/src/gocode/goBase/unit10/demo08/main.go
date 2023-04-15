package main

import "fmt"

type SayHello interface {
	sayHello()
}

type Chinese struct {
	Name string
}

type America struct {
	Name string
}

func (c Chinese) sayHello() {
	fmt.Println("你好")
}

func (a America) sayHello() {
	fmt.Println("Hello")
}

func (c Chinese) niuYangGe() {
	fmt.Println("扭秧歌")
}

func greet(s SayHello) {
	s.sayHello()

	/* 断言 */
	//ch, flag := s.(Chinese) // 看 s 是否能转成 Chinese 类型并转成 ch 变量，flag 判断返回是否转成功
	//if flag {
	//	ch.niuYangGe()
	//} else {
	//	fmt.Println("美国人不会扭秧歌")
	//}

	// 更简洁的语法
	//if ch, flag := s.(Chinese); flag {
	//	ch.niuYangGe()
	//} else {
	//	fmt.Println("美国人不会扭秧歌")
	//}

	// switch 写法
	switch s.(type) { // type 属于 go 中的一个关键字，固定写法
	case Chinese:
		ch := s.(Chinese)
		ch.niuYangGe()
	case America:
		am := s.(America)
		am.sayHello()
	}
}

func main() {
	// 定义一个 SayHello 接口数组，里面存放结构体变量
	var arr [3]SayHello
	arr[0] = America{"York"}
	arr[1] = Chinese{"YanMeng"}
	arr[2] = America{"Nan"}

	greet(arr[1])

	fmt.Println(arr)
	arr[1].sayHello()
}
