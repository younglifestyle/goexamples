package main

import (
	"fmt"
	"net"
)

func main() {
	// 指定 服务器 IP + port 创建 通信套接字。
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	fmt.Println("1")
	// 主动写数据给服务器
	conn.Write([]byte(`{""}`))
	fmt.Println("2")

	buf := make([]byte, 4096)
	// 接收服务器回发的数据
	n, err := conn.Read(buf)
	fmt.Println("3")
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	fmt.Println("服务器回发：", string(buf[:n]))
}
