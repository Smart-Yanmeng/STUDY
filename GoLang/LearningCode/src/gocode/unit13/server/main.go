package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 连接用完一定要关闭
	defer conn.Close()

	for {
		// 创建一个切片，准备将读取的数据放入切片
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端已断开连接")

			return
		}

		// 将读取内容在服务器端输出
		fmt.Println(string(buf[0:n]))
	}
}

func main() {
	// 打印
	fmt.Println("服务器启动......")

	// 监听，需要指定服务端 TCP 协议，服务端的 IP + PORT
	listen, err1 := net.Listen("tcp", "127.0.0.1:8888")
	if err1 != nil {
		fmt.Println("监听失败，具体的错误为：", err1)

		return
	}

	// 监听成功
	// 循环等待客户端的连接
	for {
		conn, err2 := listen.Accept()
		// 连接失败
		if err2 != nil {
			fmt.Println("连接失败，具体的错误为：", err2)

			return
		} else {
			// 连接成功
			fmt.Println("连接成功，具体的连接为：", conn, "接收到的客户端信息是：", conn.RemoteAddr().String())
		}

		// 准备一个协程，协程处理客户端的服务请求
		go process(conn)
	}
}
