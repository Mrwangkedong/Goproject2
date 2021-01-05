package main

import (
	"fmt"
	"time"
)

/*
通道容量为0，len=0.不能存储数据，具有同步能力，读、写同步
channel应用于两个go程中
*/

func NoneBuffer() {
	ch := make(chan int)
	//子go程
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子go程--", i)
			ch <- i
		}
	}()
	time.Sleep(2 * time.Second) //【子go程-- 0  sleep 2seconds 后续】
	//主go程
	for i := 0; i < 5; i++ {
		num := <-ch
		fmt.Println("主go程--", num)
	}

}

func main() {
	NoneBuffer()
}
