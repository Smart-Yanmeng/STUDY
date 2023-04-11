package main

import "fmt"

// Person 结构体
type Person struct {
	Name string
}

// 给 Person 结构体绑定方法 test
func (p Person) test() {
	fmt.Println(p)
}

func (p *Person) changeName() {
	// 值拷贝
	(*p).Name = "张三"
}

func main() {
	// 创建一个结构体对象
	var p Person
	p.Name = "York"

	p.test()
	p.changeName()

	fmt.Println(p)
}
