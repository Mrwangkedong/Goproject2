package main

import "fmt"

/*
死锁：
	1.单go程自己死锁：
		channel应该在至少两个以上的go程中进行通信，否则死锁
	2.go程间channel访问顺序导致死锁
	3.多go程，多channel导致死锁
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

func main() {
	ch := make(chan int)

	go func() {
		ch <- 520
	}()

	num := <-ch

	fmt.Println("num = ", num)

}
