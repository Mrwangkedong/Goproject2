package main

import (
	"fmt"
	"time"
)

/*
	Ticker
*/

func main() {
	//利用channel来控制主子go程
	ch := make(chan bool)
	i := 0
	fmt.Println("laterTime: ", time.Now())
	myTicker := time.NewTicker(time.Second * 1)
	go func() {
		for {
			nowTime := <-myTicker.C //每每获取一次，就是一次定时器，就算隔了一万年在获取，也是与之前的相差一秒
			fmt.Println("nowTime: ", nowTime)
			i++
			if i == 10 {
				ch <- true
			}
		}
	}()

	<-ch
}
