package main

import (
	"fmt"
	"time"
)

func test() {

}

func main() {
	// 启动一个协程
	// 使用了匿名函数，直接调用了匿名函数
	//go func() {
	//	fmt.Println(1)
	//}()

	// 匿名函数 + 外部变量 = 闭包
	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}

	time.Sleep(time.Second * 2)
}
