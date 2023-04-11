package main

import (
	"fmt"
)

func main() {
	// 定义一个字符串
	var str string = "Hello GoLang!"

	// 普通 for
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}

	fmt.Println("-------------------------------------")

	// for range
	for j, value := range str {
		fmt.Printf("%d: %c\n", j, value)
	}
}
