package main

import (
	"fmt"
	"net"
)

func main() {
	//1.绑定服务器的IP+port
	conn, err := net.Dial("udp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("CLient Dial err: ", err)
		return
	}
	//2.向服务器发送信息
	_, err = conn.Write([]byte("yes!!"))
	if err != nil {
		fmt.Println("CLient Write err: ", err)
		return
	}

	//接收服务器发来的信息
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)

	//展示收到的信息
	fmt.Println(string(buf[:n]))

}
