package main

import (
	"fmt"
	"os"
)

func main() {
	// 在下面的程序中，不需要进行 open 和 close 操作，因为已经被封装在 ReadFile 的内部了
	// 读取文件
	content, err := os.ReadFile("D:\\STUDY\\GoLang\\LearningCode\\src\\gocode\\unit11\\Test.txt")
	if err != nil {
		fmt.Println("读取出错，错误为：", err)
	} else {
		// 读取成功，输出内容
		fmt.Println(string(content))
	}
}
