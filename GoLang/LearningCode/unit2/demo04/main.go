package main

import "fmt"

func main() {
	// 定义字符类型的数据 - 本质上就是一个整数，可以直接参与运算，对应的是 ASCII 进行存储
	var c1 byte = 'a'
	fmt.Println(c1)

	var c2 byte = '6'
	fmt.Println(c2)

	var c3 byte = '('
	fmt.Println(c3)

	// 汉字字符，对应的是 Unicode 码值
	var c4 int = '中'
	fmt.Println(c4)
}
