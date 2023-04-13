package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 100)

	for i := 0; i < 100; i++ {
		intChan <- i
	}

	// 如果在遍历前，如果没有关闭管道，就会出现 deadlock 的错误
	// 所以我们在遍历前要关闭管道
	close(intChan)

	// 遍历
	for value := range intChan {
		fmt.Println(value)
	}
}
