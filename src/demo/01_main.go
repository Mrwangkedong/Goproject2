package main

import "fmt"

//调用结束后就会释放
func PonitTest(m int) {
	a := 100

	var p *int = &a

	fmt.Println("P指代的值 = ,P的地址 = ", *p, p)
}

func main() {

	var name string = "王柯栋"

	_ = name
	//fmt.Println(name)

	var a int = 8

	fmt.Println("a = ", a)

	var p *int = &a

	fmt.Println("*p = ", *p)

	PonitTest(88)

	//多变量声明
	//var(
	//	Myname string
	//	age int
	//	gender string
	//)
	//
	////声明多个变量
	//boy,girl := "王柯栋","王冰冰"

}
