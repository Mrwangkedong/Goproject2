package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

func main() {
	//0.指定接收文件默认文件夹
	DefaultDir := "G://GolangFileText/"
	//1.指定服务器，通信协议，IP地址，port。创建一个用于监听的socket
	myListener, err := net.Listen("tcp", "localhost:8004")
	if err != nil {
		fmt.Println("net Listener err: ", err)
		return
	}
	fmt.Println("服务端等待客户端建立连接")
	//循环监听
	for true {
		//2.阻塞监听客户端连接请求,获取客户端连接信号.若成功，返回用于通信的socket
		conn, err := myListener.Accept()
		if err != nil {
			fmt.Println("listener.Accept() err: ", err)
			return
		}
		fmt.Println("服务器与客户端成功建立连接")
		//3.利用go程处理client请求
		go HandleClient(conn, DefaultDir)
	}

}

//处理文件接收
func HandleClient(conn net.Conn, dir string) {
	//1.读取客户端发送的文件名
	buf := make([]byte, 4096)
	n, err := conn.Read(buf) //返回读到的个数
	if err != nil {
		fmt.Println("conn.Read(buf) err: ", err)
		return
	}
	//2.同意文件发送
	_, err = conn.Write([]byte("ok"))
	if err != nil {
		fmt.Println("conn.Write(buf) err: ", err)
		return
	}
	//3.接收文件
	//3.1.创建即将要接收的文件
	fwriter, err := os.Create(dir + "01" + string(buf[:n]))
	if err != nil {
		fmt.Println("Server os.Create() err: ", err)
		return
	}
	for true {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Server conn.Read(buf) err: ", err)
			return
		}
		//写入文件
		_, err = fwriter.Write(buf[:n])
		if err != nil {
			fmt.Println("Server fwriter.Write(buf[:n]) err: ", err)
			runtime.Goexit() //关闭当前go程
		}
	}
}
