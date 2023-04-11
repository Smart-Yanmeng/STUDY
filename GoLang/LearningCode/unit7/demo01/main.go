package main

import "fmt"

func main() {
	test()

	fmt.Println("异常之后的方法已被执行......")
}

func test() {
	// 利用 defer 和 recover 来捕获错误，defer 后加上匿名函数的调用
	defer func() {
		// 调用 recover 内置函数，可以捕获错误：
		err := recover()

		// 如果没有捕获到错误，返回值为零值：nil
		if err != nil {
			fmt.Println("错误已经被捕获")
			fmt.Println("err 是：", err)
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2

	fmt.Println(result)
}
