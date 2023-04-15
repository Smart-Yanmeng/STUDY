package main

import "fmt"

func main() {
	// 1. 变量的声明
	var age int

	// 2. 变量的赋值
	age = 10

	// 3. 变量的使用
	fmt.Println("age = ", age)

	// 声明和赋值可以合成一句
	var age2 int = 18
	fmt.Println("age2 = ", age2)
}
