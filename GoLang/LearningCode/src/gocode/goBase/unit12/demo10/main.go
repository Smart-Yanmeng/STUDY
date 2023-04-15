package main

import "fmt"

func main() {
	/* 默认情况下，管道是双向的 */
	// var intChan1 chan int

	/* 声明为只写 */
	var intChan2 chan<- int
	intChan2 = make(chan int, 3)

	intChan2 <- 20
	fmt.Println(intChan2)

	// 直接报错
	// n1 := <- intChan2

	/* 声明为只读 */
	var intChan3 <-chan int
	if intChan3 != nil {
		n2 := <-intChan3
		fmt.Println(n2)
	}

	// 直接报错
	// intChan3<- 30
}
