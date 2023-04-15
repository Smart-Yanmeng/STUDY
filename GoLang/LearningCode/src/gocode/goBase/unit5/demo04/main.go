package main

import "fmt"

// 定义一个求和函数
// getSum 函数返回一个函数，这个函数的参数是一个 int 类型的参数，返回值也是 int 类型
// 闭包：闭包的本质依旧是一个匿名函数，只是这个函数引入外界的变量 / 参数，匿名函数+引用的变量 / 参数 = 闭包
func getSum() func(int) int {
	sum := 0

	return func(num int) int {
		sum += num

		return sum
	}
}

func main() {
	// 闭包返回的是一个匿名函数，但是这个匿名函数引用到函数外的变量 / 参数，因此这个匿名函数就和变量 / 参数形成一个整体，构成闭包
	// 闭包中使用的变量 / 参数会一直保存在内存中，所以会一直使用 -> 意味着闭包不可滥用
	f := getSum()

	fmt.Println(f(1)) // 1
	fmt.Println(f(2)) // 3
	fmt.Println(f(3)) // 6
	fmt.Println(f(4)) // 10
}
