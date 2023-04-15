package main

import (
	"fmt"
	"time"
)

// 输出数字
func printNum() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// 做除法操作
func devide() {
	// 错误处理
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("devide() 函数出现错误，错误为：", err)
		}
	}()

	num1 := 10
	num2 := 0
	result := num1 / num2

	fmt.Println(result)
}

func main() {
	// 启动两个协程
	go printNum()
	go devide()

	time.Sleep(time.Second * 5)
}
