package main

import "fmt"

func main() {
	var age int = 18
	// &符号 + 变量，就可以获取这个变量内存的地址
	fmt.Println(&age)

	// 定义一个指针变量
	// 指向 int 类型的指针
	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr 指针的地址是", &ptr)

	// 获取 ptr 这个指针或者这个地址指向的那个数据
	fmt.Println(*ptr)
}
