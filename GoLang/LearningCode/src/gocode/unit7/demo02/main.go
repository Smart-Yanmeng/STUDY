package main

import (
	"errors"
	"fmt"
)

func main() {
	err := test()
	if err != nil {
		fmt.Println("程序中的错误是：", err)
		panic(err)
	}

	fmt.Println("异常之后的方法已被执行......")
}

func test() (err error) {
	num1 := 10
	num2 := 0

	if num2 == 0 {
		// 抛出自定义错误
		return errors.New("除数不能为 0")
	} else {
		result := num1 / num2
		fmt.Println(result)

		// 如果没有错误，则返回零值
		return nil
	}
}
