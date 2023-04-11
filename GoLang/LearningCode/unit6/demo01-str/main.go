package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	/* 统计字符串的长度，按字节进行统计 */
	str := "GoLang, 你好"   // 在 GoLang 中，汉字是 utf-8 字符集，一个汉字 3 个字节
	fmt.Println(len(str)) // 14

	/* 字符串的遍历 */
	// 方式一：for-range 键值循环
	for i, value := range str {
		fmt.Printf("索引为：%d, 具体的值为：%c\n", i, value)
	}

	// 方式二：利用 r := []rune(str)
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c ", r[i])
	}

	/* 字符串转整数 */
	num1, err := strconv.Atoi("666")
	fmt.Println(num1, err)

	/* 整数转字符串 */
	numStr := strconv.Itoa(888)
	fmt.Println(numStr)

	/* 统计一个字符串中有几个指定的子串 */
	num2 := strings.Count("java, ssm, springboot", "a")
	fmt.Println(num2)

	/* 不区分大小写的字符串比较 */
	equalFlag := strings.EqualFold("go", "GO")
	fmt.Println(equalFlag)

	/* 返回子串在字符串第一次出现的索引值，如果没有返回 -1 */
	index := strings.Index("golang", "go")
	fmt.Println(index)

	/* 字符串的替换 */
	// -1 表示全部替换
	replace := strings.Replace("goAndJava", "go", "GoLang", -1)
	fmt.Println(replace)

	/* 按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组 */
	strArr := strings.Split("go-python-java", "-")
	for i, value := range strArr {
		fmt.Printf("索引为：%d, 具体的值为：%s\n", i, value)
	}

	/* 将字符串的字母进行大小写切换 */
	fmt.Println("ToLower:", strings.ToLower("GOLANG"))
	fmt.Println("ToUpper:", strings.ToUpper("golang"))

	/* 将字符串左右两边空格去掉 */
	fmt.Println("TrimSpace:", strings.TrimSpace("     Go and Java     "))

	/* 将字符串左右两边指定字符串去掉 */
	// 去除全部
	fmt.Println("Trim:", strings.Trim("~GoLang~", "~"))

	// 去除左边
	fmt.Println("TrimLeft:", strings.TrimLeft("~GoLang~", "~"))

	// 去除右边
	fmt.Println("TrimRight:", strings.TrimRight("~GoLang~", "~"))

	/* 判断是否以指定字符串开头和结尾 */
	// 开头
	fmt.Println(strings.HasPrefix("https://java.sun.com", "https"))

	// 结尾
	fmt.Println(strings.HasSuffix("demo.png", ".png"))
}
