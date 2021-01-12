package main

import (
	"fmt"
	"net"
)

func main() {
	//指定服务器的 IP+port创建，创建套接字socket
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("客户端连接失败")
	}
	//主动写数据给服务器
	_, err = conn.Write([]byte("Are you ok?"))
	if err != nil {
		fmt.Println("客户端写入失败，err: ", err)
	}
	fmt.Println("写入成功...")

	//接收服务器发回的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("客户端接收服务端信息失败，err: ", err)
	}
	fmt.Println("客户端接收服务端信息: ", string(buf[:n]))

	err = conn.Close()
	if err != nil {
		fmt.Println("conn.Close() err: ", err)
		return
	}
}
