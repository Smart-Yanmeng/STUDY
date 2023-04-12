package main

import "fmt"

func main() {
	// 1. 定义一个字符串
	var s1 string = "Hello, GoLang!"
	fmt.Println(s1)

	// 2. 字符串是不可变的：指的是字符串一旦定义好，其中的字符的值不能改变
	var s2 string = "abc"
	s2 = "def"
	// 不可改变
	// s2[0] = 'a'
	fmt.Println(s2)

	// 3.字符串的表示形式
	// (1) 如果字符串中没有特殊字符，字符串的表示形式用双引号
	var s3 string = "aaa"
	// (2) 如果字符串中有特殊字符，字符串的表示形式用反引号
	var s4 string = `a"a"a`
	fmt.Println(s3)
	fmt.Println(s4)
}
