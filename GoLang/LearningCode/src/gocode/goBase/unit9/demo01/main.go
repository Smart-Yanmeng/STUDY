package main

import "fmt"

func main() {
	// map 的第一种创建方式
	// 定义 map 变量
	var a map[int]string
	// 只声明 map 是没有分配内存空间的
	// 必须通过 make 函数进行初始化，才会分配空间

	a = make(map[int]string, 10) // map 可以存放 10 个键值对

	// 将键值对存入 map 中
	a[20095452] = "张三"
	a[20095387] = "李四"
	a[20095777] = "王五"
	a[20095777] = "朱六"

	// map 的 key-value 是无序的
	// key 是不可以重复的，如果遇到重复，后一个 value 会替换前一个 value
	// value 可以重复
	fmt.Println(a) // map[20095387:李四 20095452:张三 20095777:朱六]

	// map 的第二种创建方式
	b := make(map[int]string)
	b[1] = "A"
	b[2] = "B"
	fmt.Println(b)

	// map 的第三种创建方式
	c := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}
	fmt.Println(c)

	// 向 map 中新增元素
	c[4] = "D"
	fmt.Println(c)

	// 向 map 中删除元素
	delete(c, 2)
	fmt.Println(c)

	// 向 map 中修改元素
	c[1] = "a"
	fmt.Println(c)

	// 清空 map - golang 中没有清空的函数，可以遍历 key，逐个删除；或者 make 一个新的 map，让原来的成为垃圾，被 gc 回收

	// 向 map 中查找元素
	value, flag := c[3]
	fmt.Println(value, flag)
}
