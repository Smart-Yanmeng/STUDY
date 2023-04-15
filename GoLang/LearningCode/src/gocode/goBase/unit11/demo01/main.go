package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("D:\\STUDY\\GoLang\\LearningCode\\src\\gocode\\unit11\\Test.txt")
	if err != nil {
		// 如果出错，输出错误信息
		fmt.Println("文件打开错误，对应错误是：", err)
	} else {
		// 没有出错，输出文件
		// file 是一个地址
		fmt.Println("文件：", file)

		// 关闭文件
		err := file.Close()

		// 关闭失败
		if err != nil {
			fmt.Println("关闭文件出错，对应的错误为：", err)
		}
	}

}
