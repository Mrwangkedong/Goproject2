package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//指定服务器的 IP+port创建，创建套接字socket
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("客户端连接失败")
	}
	defer conn.Close()

	//（子go程）从键盘中读取输入（stdin），将输入数据发送给服务器
	go func() {
		str := make([]byte, 4096)
		for true {
			//读取键盘输入
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("os.Stdin.Read(str)错误，err: ", err)
				return
			}
			//主动写数据给服务器
			_, err = conn.Write(str[:n])
			if err != nil {
				fmt.Println("客户端写入失败，err: ", err)
				continue
			}
			fmt.Println("写入成功...")

		}
	}()
	//（主go程）回显服务器回发的大写数据
	for true {
		str := make([]byte, 4096)
		//读取服务器发送的数据
		n, err := conn.Read(str)
		//服务器关闭，数据错误
		if n == 0 {
			fmt.Println("服务端关闭，客户端即将退出...")
		}
		if err != nil {
			fmt.Println("os.Stdin.Read(str)错误，err: ", err)
			return
		}
		//输出从服务器端读到的数据
		fmt.Println("从服务器端读到的数据: ", str[:n])
	}

}
