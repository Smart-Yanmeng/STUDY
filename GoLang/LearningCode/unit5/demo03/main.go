package main

import (
	"fmt"
)

// 定义一个函数
func test1(num int) {
	fmt.Println("test1", num)
}

// 定义一个函数，把另一个函数作为形参传入
func test2(num1 int, num2 float32, testFunc func(int)) {
	println(num1)
	println(num2)
	testFunc(1)
}

// 定义一个函数，求和与差
func test3(num1 int, num2 int) (sum int, sub int) {
	// 对函数返回值进行命名
	sum = num1 + num2
	sub = num1 - num2

	return
}

func main() {
	// 函数是一个类型，也可以赋值给一个变量
	a := test1
	fmt.Printf("变量 a 对应的类型是：%T\n", a)

	// 通过该变量可以对函数调用
	a(10)

	// 调用 test2 函数
	test2(1, 2.1, a)

	// 调用 test3 函数
	sum, sub := test3(1, 2)
	fmt.Println("sum = ", sum)
	fmt.Println("sub = ", sub)
}
