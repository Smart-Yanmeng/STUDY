package main

import (
	"fmt"
	"sync"
)

// 线程计数器
var wg sync.WaitGroup

// 锁
var lock sync.RWMutex

var intChan = make(chan int, 50)

func write() {
	defer wg.Done()

	for i := 0; i < 50; i++ {
		lock.Lock()
		intChan <- i
		lock.Unlock()
	}

	// 关闭管道，防止出现 deadLock 的情况
	close(intChan)
}

func read() {
	defer wg.Done()

	// 遍历管道
	for value := range intChan {
		lock.RLock()
		fmt.Println(value)
		lock.RUnlock()
	}
}

func main() {
	// 添加两个协程
	wg.Add(2)

	go write()
	go read()

	// 等待协程
	wg.Wait()
}
