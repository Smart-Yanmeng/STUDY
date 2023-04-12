package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 写入文件
	// 打开文件
	file, err := os.OpenFile("D:\\STUDY\\GoLang\\LearningCode\\src\\gocode\\unit11\\Test.txt", os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("文件打开失败")

		return
	}

	// 及时将文件关闭
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件出错")
		}
	}(file)

	// 写入文件
	writer := bufio.NewWriter(file)
	writeString, err2 := writer.WriteString("你好，GoLang")
	if err2 != nil {
		fmt.Println("写入文件出错，错误为：", err2)

		return
	}

	fmt.Println(writeString)

	// 流带缓冲区，刷新数据，真正写入文件中去
	err3 := writer.Flush()
	if err3 != nil {
		fmt.Println("刷新缓存出错，错误为：", err3)

		return
	}

	s := os.FileMode(0666).String()
	fmt.Println("当前的权限为：", s)
}
