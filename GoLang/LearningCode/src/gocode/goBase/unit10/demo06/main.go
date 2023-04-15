package main

import "fmt"

// Animal - 定义一个动物的结构体
type Animal struct {
	age    int
	weight float32
}

// Shout - 给 Animal 绑定的方法：叫
func (a *Animal) Shout() {
	fmt.Println("我可以大声喊叫")
}

func (a *Animal) ShowInfo() {
	fmt.Println("动物的年龄是：", a.age)
	fmt.Println("动物的体重是：", a.weight)
}

// Cat - 定义一个猫的结构体
type Cat struct {
	// 为了复用性，体现继承思维，加入匿名结构体
	// 将 Animal 中的字段和方法都拿到了
	Animal
}

// Scratch - 对 Cat 绑定特有的方法
func (c *Cat) Scratch() {
	fmt.Println("我是小猫，我可以挠人")
}

func main() {
	// 创建 Cat 结构体的实例
	cat := &Cat{}

	// 共有的成员变量
	cat.age = 3
	cat.weight = 10.6

	// 共有的方法
	cat.Shout()
	cat.ShowInfo()

	// Cat 自己独有的方法
	cat.Scratch()
}
