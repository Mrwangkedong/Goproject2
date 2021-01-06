package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
Select
	作用：监听channel中的数据流动
	select与switch的语法类似，由select选择一个新的选择块，每个选择条件由case语句来描述
	与switch语句相比，select有比较多的限制，UI大的限制就是每个case语句，必须是一个IO操作

	在select中，Go语言会从头至尾评估每一个发送和接收的语句
	1.如果一种任意一条语句可以继续执行，那么就从那些可以执行的语句中任意选择一条来使用
	2.如果没有任意一条语句可以执行：
		2.1.执行default（一般不太使用default）
		2.2.被阻塞，知道至少一个通信可以继续下去
	3.select本身不带有循环监听，需要外层加一个for实现
	4.break只能跳出select中的一个case
*/

func main() {
	ch := make(chan int)    //用来进行数据通信
	quit := make(chan bool) //用来判断是否退出

	go func() {
		for i := 0; i < 4; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		quit <- true     //通知主go程退出
		runtime.Goexit() //退出当前go程
	}()

	for { //主go程读数据
		select {
		case num := <-ch:
			fmt.Println("读到num=", num) //有时候会出现【0 1 2 3 0 0】，应该是出现在未关闭channel
		case <-quit:
			return
			//break   //跳出select
		}

	}
}
