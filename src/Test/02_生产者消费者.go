package main

import (
	"fmt"
	"time"
)

/*
生产者
消费者
缓冲区: 1.减少生产者，消费者之间的耦合
       2.提高并发能力
	   3.数据缓存
*/

func consumer(ch chan int) {
	time.Sleep(1 * time.Second)
	for true {
		num := <-ch
		if num == -1 {
			break
		}
		fmt.Println("接收到num=", num)
	}
}

func producer(ch chan int) {
	for i := 0; ; i++ {
		if i == 10 { //到达10之后关闭channel
			ch <- -1
			close(ch)
			break
		}
		ch <- i //向0容量缓冲区中输出
		fmt.Println("发送i=", i)
	}
}

func main() {
	ch := make(chan int, 3)
	go producer(ch)
	consumer(ch)
}
