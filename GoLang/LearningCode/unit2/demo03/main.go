package main

import "fmt"

func main() {
	var num1 = 100
	fmt.Printf("num1 的类型是：%T\n", num1) // 获取数据类型

	var num2 = 1000000
	fmt.Printf("num2 的类型是：%T\n", num2)

	var num3 = 100.001
	fmt.Printf("num3 的类型是：%T\n", num3)

	var str = "Hello"
	fmt.Printf("str 的类型是：%T\n", str)
}
