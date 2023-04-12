package main

import "fmt"

func main() {
	// 定义数组
	var intArr [6]int = [6]int{1, 2, 3, 4, 5, 6}

	// 切片构建在数组之上：
	// 定义一个切片，名字为 slice，[]动态变化的数组长度不写，int 是切片中数据的类型，intArr 为原数组
	// [1:3] 切片 - 切出的一段片段 - 索引：从 1 开始，到 3 结束 ( 不包含3 ) - [1,3)
	var slice []int = intArr[1:3]

	// 输出数组
	fmt.Println("intArr:", intArr)

	// 输出切片
	fmt.Println("slice:", slice)

	// 输出数组的元素个数
	fmt.Println("intArr的元素个数为：", len(intArr))

	// 输出切片的元素个数
	fmt.Println("slice的元素个数为：", len(slice))

	// 输出切片的容量 ( 容量可以动态变化 )
	fmt.Println("slice的元素个数为：", cap(slice))
}
