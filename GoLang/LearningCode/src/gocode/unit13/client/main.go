package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 打印
	fmt.Println("客户端启动......")

	// 调用 Dial 函数
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("客户端连接失败，具体的错误为：", err)

		return
	}
	fmt.Println("连接成功，具体的连接是：", conn)

	// 通过客户端发送单行数据，然后退出
	// os.Stdin 代表终端标准输入
	reader := bufio.NewReader(os.Stdin)

	// 从终端读取一行用户输入信息
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("终端输入失败，具体的错误为：", err)
	}

	// 将 str 数据发送给服务器
	n, err := conn.Write([]byte(str))
	if err != nil {
		fmt.Println("连接失败，具体的错误为：", err)
	} else {
		fmt.Printf("终端数据通过客户端发送，一共发送了%v个字节的数据并退出\n", n)
	}
}
