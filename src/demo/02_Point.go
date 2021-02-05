package main

//指针
import "fmt"

/**
Go Vs  C
默认值：nil   vs  NULL
操作符: & 进行取地址变量；*通过指针访问目标对象
不支持指针运算，不支持—>操作，直接用.进行访问
使用结束后要将地址置为nil
*/

/*
指针使用注意：
1、空指针，未被初始化的指针。    var p *int  ---->nil   出现error
2、野指针，被一片无效的地址空间初始化，不可以直接赋值指针
*/

/*
new
在heap（用户空间）上申请一片内存地址空间
new()  p = new(int32)
申请到的0xc0代表八位
*/

/*
变量存储
	等号左边的变量  代表  变量所指向的内存空间
	得好右边的变量  代表  变量内存空间存储的数据值
*/

/*
指针的函数传参（运行机制参照C）
	传递至（引用）：将形参的地址值作为函数参数，返回值后传递
	传值（数据）：将形参的值拷贝一份给实参
*/

//使用new进行新指针的申请
func creatNewPonint() {
	var p *int32
	var q *int64
	p = new(int32)
	q = new(int64)
	*p = 32
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(q)
}

//利用指针进行值的交换
//func swapPoint(p,q *int)  简化
func swapPoint(p *int, q *int) {
	*p, *q = *q, *p
}

func main() {
	creatNewPonint()

	//值的交换
	var A int = 100
	var B int = 200
	swapPoint(&A, &B)
	fmt.Println(A, B)

	//var name int = 123
	//fmt.Println("123")
	//fmt.Printf("%d",name)
}
