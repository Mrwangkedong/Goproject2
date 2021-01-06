package main

import "fmt"

// 1 1 2 3 5 8 13 21

func fibonacci(ch <-chan int, quit <-chan bool) {
	for true {
		select {
		case num := <-ch:
			fmt.Printf("%d\t", num)
		case <-quit:
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go fibonacci(ch, quit)
	//主go程
	x, y := 1, 1
	ch <- x
	ch <- y
	for i := 0; i < 15; i++ {
		x, y = y, x+y
		ch <- y //向通道中写入
	}
	quit <- true

}
