package main

import (
	"fmt"
	"time"
)

/*
channel是go语言中的一个核心类型，可以把它看做"管道"，并发核心单元通过它就可以发送或者接受数据进行通讯，在一定程度上进一步降低了编程的难度
channel是一个数据类型，主要用来解决go程的同步问题以及协程之间的数据共享（数据传递）问题
goroutine运行在相同的地址空间，因此访问共享内存必须做好同步，goroutine奉行“通过通信来共享内存，而不是共享内存来通信”
goroutine同步是在channel里面完成的
channel,一端进行读，一端进行写
【补充知识点】：每当有一个进程启动时，系统会自动打卡三个文件：标准输入，标准输出，标准错误----stdin/stdout/stderr
*/

/*
写入端：channel <- 8
写出端  <- channel
要求读端和写端必须同时满足条件，才在
*/

func creatChannel() {
	chanl := make(chan int) //make(chan 在channel中传递的数据类型，容量)  容量==0，无缓冲channel，容量>0,有缓冲channel
	_ = chanl
}

//全局定义channel,用来完成数据通信
var channel = make(chan int)

//定义一台打印机
func printer(s string) {
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(300 * time.Millisecond)
	}
}

//定义两个人使用打印机
func person1() { //先执行
	printer("Hello")
	channel <- 8 //将8写入channel中去，只有后续将他读出来才会解锁！！！
}
func person2() { //后执行
	<-channel //将channel中的数据取出来，相当于解锁
	printer("world")
}

func channelLock() {
	////输出【Hweorllldo】
	//go person1()
	//go person2()
	//for true {
	//	;
	//}
	//输出【helloword】
	go person1()
	go person2()
	for true {

	}
}

/*

 */
func channelSameStep() {
	ch := make(chan string)
	//len(ch)表示通道中未被读取的数据个数，cap(ch)表示通道容量
	println("len(ch)=", len(ch), "--->cap(ch)=", cap(ch)) //一直是0
	//子 go程
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d", i)
		}
		//通知子go程打印完毕
		ch <- "子go程打印结束"
	}()
	println("len(ch)=", len(ch), "--->cap(ch)=", cap(ch)) //一直是0
	//主 go程,接收通道内容
	str := <-ch
	println("\n" + str)
}

func main() {
	channelSameStep()

}
