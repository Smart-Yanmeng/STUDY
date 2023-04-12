package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("D:\\STUDY\\GoLang\\LearningCode\\src\\gocode\\unit11\\Test.txt")
	if err != nil {
		fmt.Println("文件打开失败，错误为：", err)
	}

	// 当函数退出时，让 file 关闭，防止内存泄漏
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件出错")
		}
	}(file)

	// 创建一个流
	reader := bufio.NewReader(file)

	// 读取操作
	for {
		// 读取到一个换行就结束
		readString, err := reader.ReadString('\n')

		// 当读取到 io.EOF 的时候，就说明已经读取到文件的结尾了
		if err == io.EOF {
			break
		}

		// 如果没有读取到文件的末尾时，输出文件内容
		fmt.Println(readString)
	}

	// 结束
	fmt.Println("文件读取成功")
}
