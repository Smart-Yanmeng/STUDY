package main

import "fmt"

/* 全局变量 */
var n7 = 100
var n8 = 101

// 一次性声明
var (
	n9  = "NAN"
	n10 = 200
)

func main() {
	/* 局部变量 */
	// 1. 变量的使用方式
	var num1 int = 18
	fmt.Println(num1)

	// 2. 变量有默认值
	var num2 int
	fmt.Println(num2)

	// 3. 自动类型推断
	var num3 = 10
	fmt.Println(num3)

	// 4. 省略 var 关键字
	sex := "男"
	fmt.Println(sex)

	fmt.Println("-----------------------")

	// 声明多个变量
	var n1, n2, n3 int
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	n4, n5, n6 := "Hello", 1, 2.3
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)

	fmt.Println("-----------------------")
	fmt.Println(n7)
	fmt.Println(n8)

	fmt.Println(n9)
	fmt.Println(n10)
}
