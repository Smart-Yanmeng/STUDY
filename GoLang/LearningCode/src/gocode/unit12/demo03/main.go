package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // 只定义，无需赋值

func main() {
	for i := 0; i < 5; i++ {
		// 协程开始的时候加 1 操作
		// Add 中加入的数字要和协程的次数保持一致
		wg.Add(1)
		go func(n int) {
			// 主线程一直在阻塞，什么时候 wg 减为 0 了，就停止
			// 防止忘记计数器，就使用 defer 关键字
			defer wg.Wait()

			fmt.Println(n)
			// 协程结束的时候减 1 操作
			wg.Done()
		}(i)
	}

	//wg.Wait()
}
