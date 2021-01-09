package main

import "sync"

/*
条件变量：
	本身不是锁！！！但经常与锁结合使用
sync.Cond类型表示了条件变量。成员变量L代表与条件变量搭配使用的锁
常用方法有三个：Wait(阻塞)、Signal(释放)、Broadcast
*/

/*
func (c *Cond) wait():（被阻塞，进入阻塞队列。就释放当前已经掌握的资源）
	1.阻塞等待条件变量满足
	2.释放已经掌握的互斥锁，相当于cond.L.Unlock()           注意：两步为一个原子操作
	3.当被唤醒，Wait()函数返回时，解除阻塞并重新获取互斥锁，相当于cond.L.Lock()
*/
/*
func (c *Cond) Signal()
	单发通知，给一个正等待（阻塞）在该条件变量上的go程发送通知（一个）
func (c *Cond) Broadcast()
	广播通知，给正在等待（阻塞）在该条件不练上的所有go程发送通知（全部）
*/

var myCond sync.Cond

func main() {

}
