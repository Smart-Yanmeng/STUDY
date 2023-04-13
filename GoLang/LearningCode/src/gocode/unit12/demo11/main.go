package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个 int 管道
	intChan := make(chan int, 1)

	go func() {
		time.Sleep(time.Second * 5)
		intChan <- 10
	}()

	// 定义一个 string 管道
	stringChan := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		stringChan <- "Hello"
	}()

	// 本身取数据就是阻塞的
	// fmt.Println(<-intChan)

	// 解决多个管道的选择问题，也可以叫做多路复用，可以从多个管道中随机公平地选择一个来执行
	select {
	case v := <-intChan:
		fmt.Println("intChan:", v)
	case v := <-stringChan:
		fmt.Println("stringChan:", v)
	default:
		fmt.Println("防止 select 被阻塞")
	}
}
