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

/*
channel关闭：
		如果发送者知道，没有更多的值需要发送到channel的话，那么让接受者也能即及时知道没有多余的值可接收将是有用的，
  	因为接收者可以停止不必要的接收等待。这可以通过内置的clise函数来关闭channel实现。
		已经关闭的channel不能进行写入，可以再读，但是再接收得到话只能读到0 or default
*/

func fullChannel() {
	ch := make(chan int, 3) //存满三个元素之前，不会阻塞
	fmt.Println("len(ch)", len(ch), "cap(ch)", cap(ch))
	//子go程
	go func() {
		for i := 1; i < 9; i++ {
			ch <- i //先写入，再打印
			fmt.Println("子go程，i=", i, "len(ch)", len(ch), "cap(ch)", cap(ch))
			if i == 6 {
				close(ch) //关闭channel
			}
		}
	}()
	time.Sleep(3 * time.Second)
	//主go程
	for i := 1; i < 9; i++ {
		//判断channel是否关闭,关闭返回false
		if num, ok := <-ch; ok {
			fmt.Println("主go程，num=", num, "len(ch)", len(ch), "cap(ch)", cap(ch))
		} else {
			fmt.Println("channel已关闭...")
		}
		//num := <-ch
		//<- ch		//丢弃

	}
}

func main() {
	fullChannel()
}
