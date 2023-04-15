package main

import "fmt"

func main() {
	// 管道的本质是一个队列
	// 定义一个 int 类型的管道
	var intChan chan int

	// 通过 make 初始化：管道可以存放 3 个 int 类型的数据
	intChan = make(chan int, 3)

	// 证明管道是引用类型
	fmt.Printf("intChan 的值：%v\n", intChan)

	// 向管道存放数据
	intChan <- 10
	num := 20
	intChan <- num

	// 输出管道的长度
	fmt.Printf("管道的实际长度是：%v，管道的容量是：%v\n", len(intChan), cap(intChan))

	// 在管道中读取数据
	num1 := <-intChan

	fmt.Println("num1 =", num1)
}
