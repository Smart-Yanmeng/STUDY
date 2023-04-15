package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义一个变量
var totalNum int

var wg sync.WaitGroup

// 加入读写锁
var lock sync.RWMutex

func read() {
	defer wg.Done()

	// 开锁 - 如果只是读数据，那么这个锁不产生影响，但是读写同时发生的时候，就会有影响
	lock.RLock()

	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取数据成功")

	// 关锁
	lock.RUnlock()
}

func write() {
	defer wg.Done()
	lock.Lock()

	fmt.Println("开始写入数据")
	time.Sleep(time.Second * 10)
	fmt.Println("写入数据成功")

	lock.Unlock()
}

func main() {
	wg.Add(6)

	// 启动协程 -> 读多写少
	for i := 0; i < 5; i++ {
		go read()
	}
	go write()

	wg.Wait()
	fmt.Println(totalNum)
}
