package main

import "fmt"

// 定义一个函数，函数的参数为可变参数
// ... 表示参数的数量是可变的
func print(args ...int) {
	// 函数内部处理可变参数的时候，将可变参数当做切片来处理
	// 遍历输出
	for i, value := range args {
		fmt.Println(i, value)
	}
}

func main() {
	print(1, 2, 3, 4, 5, 6, 7, 8, 9)
}
