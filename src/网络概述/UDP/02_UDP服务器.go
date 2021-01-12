package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//1.组织一个udp地址结构,指定服务器的IP+Port
	serAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("ResolveUDPAddr err: ", err)
		return
	}
	fmt.Println("UDP 服务器地址结构创建完成...")
	//2.创建用于通信的socket
	udpConn, err := net.ListenUDP("udp", serAddr)
	if err != nil {
		fmt.Println("ListenUDP err: ", err)
		return
	}
	fmt.Println("socket建立，等待客户端连接...")
	defer udpConn.Close()

	buf := make([]byte, 4096)
	//3.读取客户端发送的数据
	n, clientAddr, err := udpConn.ReadFromUDP(buf)
	if n == 0 {
		fmt.Println("server closed...\n client closeing...")
		return
	}
	if err != nil {
		fmt.Println("ReadFromUDP err: ", err)
		return
	}
	//4.展示读取到的数据
	fmt.Println(clientAddr, "：", string(buf[:n]))

	//5.写回数据给客户端
	_, err = udpConn.WriteToUDP([]byte(strings.ToUpper(string(buf[:n]))), clientAddr)
	if err != nil {
		fmt.Println("WriteToUDP err: ", err)
		return
	}

}
