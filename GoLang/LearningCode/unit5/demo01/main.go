package main

import "fmt"

func main() {
	// 函数调用
	num1 := 10
	num2 := 20

	result1, result2 := sumAnd(num1, num2)
	result3, _ := sumAnd(num1, num2)

	fmt.Println("和为：", result1)
	fmt.Println("积为：", result2)
	fmt.Println("我只拿到了一个和为：", result3)
}

func sumAnd(a int, b int) (int, int) {
	return a + b, a * b
}
