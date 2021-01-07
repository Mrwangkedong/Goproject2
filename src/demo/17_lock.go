package main

import (
	"fmt"
	"sync"
	"time"
)

/*
死锁：
	1.单go程自己死锁：
		channel应该在至少两个以上的go程中进行通信，否则死锁
	2.go程间channel访问顺序导致死锁
	3.多go程，多channel导致死锁
	4.在Go语言中，尽量不要把互斥锁，读写锁，与channel混用
*/

/*
互斥锁：
	10_channel.go       //利用channel完成
	func printer2       //利用Lock完成
*/

//死锁demo，channel必须要两端把持。两个go程
func DeadLockDemo01() {
	ch := make(chan int)

	ch <- 520 //造成写端阻塞

	num := <-ch

	fmt.Println("num = ", num)

}
func DeadLockDemo02() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	//子go程
	go func() {
		for true {
			select {
			case num := <-ch1:
				ch2 <- num
			}
		}
	}()
	//主go程
	for true {
		select {
		case num := <-ch2:
			ch1 <- num
		}
	}
}

//使用“锁”完成同步

//设置锁的全局变量
var mutex sync.Mutex //创建一个互斥量，新建的互斥锁状态为0，未加锁

func printer2(str string) {
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(300 * time.Millisecond)
	}
}

func people1() { //先
	mutex.Lock() //上锁
	printer2("Hello,")
	mutex.Unlock() //解锁
}
func people2() { //后
	mutex.Lock() //上锁
	printer2("World!")
	mutex.Unlock() //解锁
}

func main() {
	go people1()
	go people2()

	for true {

	}
}
