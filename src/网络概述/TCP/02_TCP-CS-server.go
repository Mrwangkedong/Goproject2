package main

import (
	"fmt"
	"net"
)

func main() {
	//指定服务器，通信协议，IP地址，port。创建一个用于监听的socket
	myListener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("net Listener err: ", err)
		return
	}
	fmt.Println("服务端等待客户端建立连接")
	//阻塞监听客户端连接请求,获取客户端连接信号.若成功，返回用于通信的socket
	conn, err := myListener.Accept()
	if err != nil {
		fmt.Println("listener.Accept() err: ", err)
		return
	}
	fmt.Println("服务器与客户端成功建立连接")
	//读取客户端发送的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf) //返回读到的个数
	if err != nil {
		fmt.Println("conn.Read(buf) err: ", err)
		return
	}
	//处理数据---打印
	fmt.Println("服务器读到的数据：", string(buf[:n]))

	//服务端给客户端写数据
	_, err = conn.Write([]byte("Yes！I am ok!"))
	if err != nil {
		fmt.Println("conn.Write(buf) err: ", err)
		return
	}

	//关闭socket
	err = myListener.Close()
	if err != nil {
		fmt.Println("myListener.Close() err: ", err)
		return
	}
	err = conn.Close()
	if err != nil {
		fmt.Println("conn.Close() err: ", err)
		return
	}

}
