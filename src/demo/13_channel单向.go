package main

import "fmt"

/*
单向channel：
	1.默认的channel是双向的
	2.单向写channel   var sendCh chan <- int    sendCh = make(chan -< int)
	  单向读channel   var recvCh chan <- int	   recvCh = make(<- chan int)
	3.单向双向转换：
		双向的可以隐身转化为任意一种单向channel   sendCh = ch
		单向channel不能转换为双向
	4.传参：传【引用】
*/

func send(out chan<- int) {
	out <- 250
	close(out)
}

func recv(in <-chan int) {
	num := <-in
	fmt.Println("读取到：", num)
}

func main() {
	ch := make(chan int)
	go func() {
		send(ch)
	}()

	recv(ch)

}
