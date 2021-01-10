package main

import (
	"fmt"
	"net"
	"runtime"
	"strings"
)

func main() {
	//创建监听套接字
	//指定服务器，通信协议，IP地址，port。创建一个用于监听的socket
	myListener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("net Listener err: ", err)
		return
	}
	defer myListener.Close()
	fmt.Println("服务端等待客户端建立连接")

	//阻塞监听客户端连接请求,获取客户端连接信号.若成功，返回用于通信的socket
	//利用循环多重复监听
	for {
		conn, err := myListener.Accept()
		if err != nil {
			fmt.Println("listener.Accept() err: ", err)
			return
		}
		fmt.Println("服务器与客户端成功建立连接")

		//封装函数，具体完成服务器和客户端的数据通信
		go HandlerConnect(conn)
	}

}

//单独处理go程连接
func HandlerConnect(conn net.Conn) {
	defer conn.Close()

	//获取客户端的地址
	addr := conn.RemoteAddr()
	fmt.Println("客户端IP：", addr)

	//读取客户端发送的数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf) //返回读到的个数
		if n == 0 {
			fmt.Println("服务器检测到客户端已经关闭，断开连接...")
		}
		if err != nil {
			fmt.Println("conn.Read(buf) err: ", err)
			return
		}

		//退出标识符【当客户端输入“exit”，服务器收到的是“exit\n”】
		if strings.HasPrefix(string(buf[:n]), "exit") == true {
			fmt.Println("客户端要求退出")
			runtime.Goexit()
		}

		//处理数据---打印
		fmt.Println("服务器读到的数据：", string(buf[:n]))

		//服务端给客户端写数据(小写转大写)
		_, err = conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
		if err != nil {
			fmt.Println("conn.Write(buf) err: ", err)
			return
		}
	}

}
