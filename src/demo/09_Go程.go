package main

/*
coroutine程创建于进程之中，直接使用go关键字，放置于函数调用前面，产生一个go程，并发
coroutine的特性：主go程结束后随之退出(很重要)
*/

/*
runtime.Gosched()出让当前go所占用的时间片，当再次获得cpu时，从出让位置继续开始执行
*/
/*
runtime.Goexit()将立即终止当前goroutine执行，调度器确保所有已注册的defer延迟调用被执行
	return:返回当前函数调用，return之前的defer会有效
	Goexit()：结束调用该函数的Go程
*/
/*
runtime.GOMAXPROCS():用来设置可以并行计算的cpu核数的最大值，并返回之前的值
*/
import (
	"fmt"
	"runtime"
	"time"
)

func sing() {
	for i := 0; i < 5; i++ {
		fmt.Println("正在唱歌....")
		time.Sleep(100 * time.Millisecond) //sleep后放弃cpu，增加其他进程抢夺概率
	}

}

func dancing() {
	for i := 0; i < 5; i++ {
		fmt.Println("正在跳舞....")
		time.Sleep(100 * time.Millisecond)
	}
}

func test() {
	defer fmt.Println("cccccccccc")
	runtime.Goexit() //退出当前Go程
	defer fmt.Println("dddddddddd")
}

//runtime.Gosched()
func GoschedTest() {
	go sing()
	go dancing()
	go func() {
		runtime.Gosched()
		for i := 0; i < 5; i++ {
			fmt.Println("正在游泳....")
			time.Sleep(100 * time.Millisecond)
		}
	}()
	select {}
}

//runtime.Goexit()
func GoexitTest() {
	go func() {
		fmt.Println("aaaaaaaa")
		//test()   //退出大的Go程
		go test() //退出test  go程
		defer fmt.Println("bbbbbbbbbbb")

	}()

	for {

	}
}

//runtime.GOMAXPROCS()
func GOMAXPROCSTest() {
	//for  {
	//	go fmt.Print(0)   //子go程                     输出：10011010110010011111000010111001001111100
	//	fmt.Print(1)	//主go程
	//}

	fmt.Println(runtime.GOMAXPROCS(1)) //输出 4【4核cpu】，将cpu设置为单核
	for {
		go fmt.Print(0) //子go程				  输出：00000000000111111111111000000000000111111111（一个cpu你来我往）
		fmt.Print(1)    //主go程
	}
}

func main() {
	//GoexitTest()
	GOMAXPROCSTest()
}
