package main

import (
	"fmt"
	"time"
)

/*
time.Timer
	Timer是一个定时器，代表未来的一个单一事件，你可以告诉Timer你要等待多长时间
	type Timer struct{
		C <- chan time    //读channel
		r runtimeTimer
	}
	  它提供一个channel，在定时时间到达之前，没有数据写入timer.C，会一直阻塞。
	知道定时时间到，系统会自动向timer.C这个channel中写入当前时间。阻塞即被解除。
*/
//三种定时方法
func ThreeTimer() {
	//1.sleep
	time.Sleep(time.Second * 1)
	//2.Timer.C
	myTimer := time.NewTimer(1 * time.Second)
	nowTimer := <-myTimer.C //定时满，系统自动写入系统时间
	_ = nowTimer
	//3.time.After
	nowTimer = <-time.After(time.Second * 1)
}

//定时器的停止，重置
func alterTimer() {
	myTimer := time.NewTimer(3 * time.Second) //创建定时器
	myTimer.Reset(1 * time.Second)            //重置定时器
	//子go程1
	go func() {
		<-myTimer.C //读端把持，写端由系统把持（Stop之后会出现死锁）
		fmt.Println("子go程 定时完毕")
	}()
	myTimer.Stop() //设置定时器停止
	//子go程2
	go func() {
		time.Sleep(4 * time.Second)
	}()

	select {}
}

func main() {
	alterTimer()
}
