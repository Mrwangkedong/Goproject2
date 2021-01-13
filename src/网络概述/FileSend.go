package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

//文件属性
func main() {
	list := os.Args //获取命令行参数，在main函数启动时，向整个程序传参

	fileLocal := list[1]
	//1.根据文件名，获取文件属性
	fileInfo, err := os.Stat(fileLocal)
	if err != nil {
		fmt.Println("os.Stat err: ", err)
	}
	fmt.Println("文件名：", fileInfo.Name())
	fmt.Println("文件大小：", fileInfo.Size())

	CheckPower(fileInfo, fileLocal)

}

func CheckPower(fileInfo os.FileInfo, fileLocal string) {
	//1.与服务器建立Dial
	conn, err := net.Dial("tcp", "127.0.0.1:8004")
	if err != nil {
		fmt.Println("Client Dial err: ", err)
		return
	}
	defer conn.Close()
	//2.发送文件名给服务器
	_, err = conn.Write([]byte(fileInfo.Name()))
	if err != nil {
		fmt.Println("Client Write err: ", err)
		return
	}
	//3.接收服务器 回执
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if n == 0 {
		fmt.Println("服务器已关闭...")
	}
	if err != nil {
		fmt.Println("Client Read err: ", err)
		return
	}
	//4.判断服务器是否同意发送,同意发送则进行文件发送
	if string(buf[:2]) == "ok" {
		SendFile(conn, fileLocal)
		fmt.Println("Client 发送完成...")
	} else {
		fmt.Println("服务器不同意文件发送...")
	}

}

//发送文件
func SendFile(conn net.Conn, fileLocal string) {
	//1.打开文件
	freader, err := os.OpenFile(fileLocal, os.O_RDWR, 6)
	if err != nil {
		fmt.Println("打开读入文件错误：: ", err)
	}
	defer freader.Close()
	//2.进行文件发送
	buf := make([]byte, 1024*4)
	for {
		//3.for循环读取全部文件内容并通过socket发送
		n, err := freader.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("读完了、、、、")
			return
		}
		//4.写到服务器
		_, err = conn.Write(buf[:n])
		if err != nil && err == io.EOF {
			fmt.Println("Client Write err: ", err)
			return
		}
	}
}
