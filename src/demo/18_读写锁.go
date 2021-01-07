package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
读写锁：
	锁还是一把，但他有了属性。
	读写锁可以让多个读操作并发，同时读取，但是对于写操作是完全互斥的。也就是说，当一个go程进行写操作时，其他go程既不讷讷个进行读，也不能进行写。
	Go中的读写锁有由结构体类型sync.RWMutex表示。此类型的方法集合中包括两对方法：
		1.对写操作的锁定和解锁【写锁定】【写解锁】
		2.对读操作的锁定和解锁【读锁定】【读解锁】
*/

//创建读写锁
var rwMutex sync.RWMutex

//读进程
func reafGo(in <-chan int, index int) {
	rwMutex.RLock() //以读模式加锁
	num := <-in
	fmt.Printf("%dth 读go程，读入，%d\n", index, num)
	rwMutex.RUnlock() //以读模式解锁
}

//写进程
func writeGo(out chan<- int, index int) {
	//生成随机数
	num := rand.Intn(1000)
	rwMutex.Lock() //以写模式加锁
	out <- num
	fmt.Printf("%dth 写go程，写入，%d\n", index, num)
	time.Sleep(1 * time.Second)
	rwMutex.Unlock() //以写模式解锁
}

func RWLock() {
	ch := make(chan int) //用于数据传递

	//播种随机数种子
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		go writeGo(ch, i+1)
	}

	for i := 0; i < 5; i++ {
		go reafGo(ch, i+1)
	}

	for true {

	}
}

func main() {
	RWLock()
}
