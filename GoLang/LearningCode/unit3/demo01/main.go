package main

import "fmt"

func main() {
	// 实现功能：键盘录入学生的年龄、姓名、成绩、是否是 VIP
	var age int
	var name string
	var score float32
	var isVIP bool

	/* 方法一 - Scanln */
	//fmt.Println("请输入学生的年龄：")
	//fmt.Scanln(&age)
	//
	//fmt.Println("请输入学生的姓名：")
	//fmt.Scanln(&name)
	//
	//fmt.Println("请输入学生的成绩：")
	//fmt.Scanln(&score)
	//
	//fmt.Println("请输入学生是否是 VIP：")
	//fmt.Scanln(&isVIP)

	/* 方法二 - Scanf */
	fmt.Println("请输入学生的年龄、姓名、成绩、是否为 VIP，用空格隔开")
	fmt.Scanf("%d %s %f %t", &age, &name, &score, &isVIP)

	// 将上述数据在控制台打印输出
	fmt.Printf("age = %v, name = %v, score = %v, isVIP = %v", age, name, score, isVIP)
}
