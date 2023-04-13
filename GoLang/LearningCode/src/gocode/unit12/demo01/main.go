package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello, GoLang + " + strconv.Itoa(i))

		// 阻塞 1 秒
		time.Sleep(time.Second)
	}
}

func main() { // 主线程 - 主死从随
	go test() // 开启一个协程

	for i := 0; i < 10; i++ {
		fmt.Println("Hello, York + " + strconv.Itoa(i))

		// 阻塞 1 秒
		time.Sleep(time.Second)
	}
}
