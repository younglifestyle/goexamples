package main

import (
	"fmt"
	"net"
)

func main() {
	// 指定 服务器 IP + port 创建 通信套接字。
	conn, err := net.Dial("udp", "127.0.0.1:8006")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 1000000; i++ {
		// 主动写数据给服务器
		conn.Write([]byte("Are you Ready?"))

		buf := make([]byte, 4096)
		// 接收服务器回发的数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("服务器回发：", string(buf[:n]))
	}

}
