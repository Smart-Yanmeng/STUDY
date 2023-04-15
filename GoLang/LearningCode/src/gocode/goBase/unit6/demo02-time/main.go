package main

import (
	"fmt"
	"time"
)

func main() {
	// Now() 返回值是一个结构体，类型是 time.Time
	fmt.Println(time.Now())

	/* 获取日期 */
	fmt.Println("Year:", time.Now().Year())
	// 默认为英文
	fmt.Println("Month:", int(time.Now().Month()))
	fmt.Println("Day:", time.Now().Day())
	fmt.Println("Hour:", time.Now().Hour())
	fmt.Println("Minute:", time.Now().Minute())
	fmt.Println("Second:", time.Now().Second())

	fmt.Println("---------------------------------")

	/* 格式化日期 */
	now := time.Now()
	fmt.Printf("当前日期：\n年月日：%d-%d-%d\n时分秒：%d:%d:%d\n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	fmt.Println("---------------------------------")

	// Sprintf 可以得到这个字符串以便后续使用
	date := fmt.Sprintf("当前日期：\n年月日：%d-%d-%d\n时分秒：%d:%d:%d\n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(date)

	fmt.Println("---------------------------------")

	// 执行格式，这个参数字符串的各个数字必须是固定的，必须这样写
	// 选择任意的组合都是可以的，根据需求自己选择就可以
	dateStr := now.Format("2006/04/02 15/04/05")
	fmt.Println(dateStr)
}
