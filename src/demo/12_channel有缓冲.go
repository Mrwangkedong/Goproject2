package main

import (
	"fmt"
	"time"
)

/*
有缓冲channel，len(ch):channel中剩余未读取数据的个数，cap(ch)：通道的容量
channel应用于两个go程中，一个读，一个写
缓冲区可以进行数据存储，存储至容量上限，阻塞。具有异步能力
*/

func fullChannel() {
	ch := make(chan int, 3) //存满三个元素之前，不会阻塞
	fmt.Println("len(ch)", len(ch), "cap(ch)", cap(ch))
	//子go程
	go func() {
		for i := 1; i < 9; i++ {
			ch <- i //先写入，再打印
			fmt.Println("子go程，i=", i, "len(ch)", len(ch), "cap(ch)", cap(ch))
		}
	}()
	time.Sleep(3 * time.Second)
	//主go程
	for i := 1; i < 9; i++ {
		num := <-ch
		fmt.Println("主go程，num=", num, "len(ch)", len(ch), "cap(ch)", cap(ch))
	}
}

func main() {
	fullChannel()
}
