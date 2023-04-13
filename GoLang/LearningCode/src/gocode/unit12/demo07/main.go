package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3)

	intChan <- 10
	intChan <- 20

	// 关闭管道
	close(intChan)

	// 再次写入数据就会报错
	// intChan <- 30

	// 当管道关闭时，读取数据是不会报错的
	num := <-intChan
	fmt.Println(num)
}
